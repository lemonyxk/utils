/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2022-05-29 09:43
**/

package utils

import (
	"archive/tar"
	zip2 "archive/zip"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
)

const Compress compress = iota

type compress int

type srcPath struct {
	absPath string
	err     error
}

func (c compress) From(path string) *srcPath {
	var asbPath, err = filepath.Abs(path)
	return &srcPath{absPath: asbPath, err: err}
}

func (s *srcPath) Zip(dst string) error {
	if s.err != nil {
		return s.err
	}

	var absPath, err = filepath.Abs(dst)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(s.absPath, filepath.Base(absPath))); err == nil {
		return errors.New(absPath + " is exists")
	}

	fStat, err := os.Stat(s.absPath)
	if err != nil {
		return err
	}

	absDstPath := filepath.Dir(absPath)
	_, err = os.Stat(absDstPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(absDstPath, 0755)
		if err != nil {
			return err
		}
	}

	// is dir
	if fStat.IsDir() {
		var files = Dir.New(s.absPath).ReadAll()

		fw, err := os.Create(absPath)
		defer func() { _ = fw.Close() }()
		if err != nil {
			return err
		}

		var zw = zip2.NewWriter(fw)
		defer func() { _ = zw.Close() }()

		for i := 0; i < len(files); i++ {
			err = files[i].LastError()
			if err != nil {
				return err
			}
			err = doZip(absDstPath, files[i].Path(), files[i].Info(), zw)
			if err != nil {
				return err
			}
		}
		return err
	} else {

		fw, err := os.Create(absPath)
		defer func() { _ = fw.Close() }()
		if err != nil {
			return err
		}

		var zw = zip2.NewWriter(fw)
		defer func() { _ = zw.Close() }()

		return doZip(absDstPath, s.absPath, fStat, zw)
	}
}

func doZip(dst, path string, fi os.FileInfo, zw *zip2.Writer) error {
	fh, err := zip2.FileInfoHeader(fi)
	if err != nil {
		return err
	}

	fh.Name = path[len(dst)+1:]

	if fi.IsDir() {
		fh.Name += string(filepath.Separator)
	}

	w, err := zw.CreateHeader(fh)
	if err != nil {
		return err
	}

	if !fh.Mode().IsRegular() {
		return nil
	}

	fr, err := os.Open(path)
	defer func() { _ = fr.Close() }()
	if err != nil {
		return err
	}
	_, err = io.Copy(w, fr)
	if err != nil {
		return err
	}

	return nil
}

func (s *srcPath) UnZip(dst string) error {

	if s.err != nil {
		return s.err
	}

	var absDstPath, err = filepath.Abs(dst)
	if err != nil {
		return err
	}

	_, err = os.Stat(absDstPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(absDstPath, 0755)
		if err != nil {
			return err
		}
	}

	cf, err := zip2.OpenReader(s.absPath)
	if err != nil {
		return err
	}
	defer func() { _ = cf.Close() }()

	return doUnzip(cf.File, absDstPath)
}

func doUnzip(cf []*zip2.File, absPath string) error {
	for _, file := range cf {

		if file.FileInfo().IsDir() {
			_ = os.Mkdir(path.Join(absPath, file.Name), 0755)
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}

		f, err := os.Create(path.Join(absPath, file.Name))
		if err != nil {
			_ = rc.Close()
			continue
		}

		_, err = io.Copy(f, rc)
		if err != nil {
			_ = rc.Close()
			_ = f.Close()
			return err
		}

		_ = rc.Close()
		_ = f.Close()
	}

	return nil
}

func (s *srcPath) TarGz(dst string) error {
	if s.err != nil {
		return s.err
	}

	var absPath, err = filepath.Abs(dst)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(s.absPath, filepath.Base(absPath))); err == nil {
		return errors.New(absPath + " is exists")
	}

	fStat, err := os.Stat(s.absPath)
	if err != nil {
		return err
	}

	absDstPath := filepath.Dir(absPath)
	_, err = os.Stat(absDstPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(absDstPath, 0755)
		if err != nil {
			return err
		}
	}

	// is dir
	if fStat.IsDir() {
		var files = Dir.New(s.absPath).ReadAll()

		fw, err := os.Create(absPath)
		defer func() { _ = fw.Close() }()
		if err != nil {
			return err
		}

		gw := gzip.NewWriter(fw)
		defer func() { _ = gw.Close() }()
		tw := tar.NewWriter(gw)
		defer func() { _ = tw.Close() }()

		for i := 0; i < len(files); i++ {
			err = files[i].LastError()
			if err != nil {
				return err
			}
			err = doTarGz(absDstPath, files[i].Path(), files[i].Info(), tw)
			if err != nil {
				return err
			}
		}
		return err
	} else {

		fw, err := os.Create(absPath)
		defer func() { _ = fw.Close() }()
		if err != nil {
			return err
		}

		gw := gzip.NewWriter(fw)
		defer func() { _ = gw.Close() }()
		tw := tar.NewWriter(gw)
		defer func() { _ = tw.Close() }()

		return doTarGz(absDstPath, s.absPath, fStat, tw)
	}
}

func doTarGz(dst, path string, fi os.FileInfo, zw *tar.Writer) error {
	fh, err := tar.FileInfoHeader(fi, fi.Name())
	if err != nil {
		return err
	}

	fh.Name = path[len(dst)+1:]

	if fi.IsDir() {
		fh.Name += string(filepath.Separator)
	}

	err = zw.WriteHeader(fh)
	if err != nil {
		return err
	}

	if !fh.FileInfo().Mode().IsRegular() {
		return nil
	}

	fr, err := os.Open(path)
	defer func() { _ = fr.Close() }()
	if err != nil {
		return err
	}
	_, err = io.Copy(zw, fr)
	if err != nil {
		return err
	}

	return nil
}

func (s *srcPath) UnTarGz(dst string) error {

	if s.err != nil {
		return s.err
	}

	var absDstPath, err = filepath.Abs(dst)
	if err != nil {
		return err
	}

	_, err = os.Stat(absDstPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(absDstPath, 0755)
		if err != nil {
			return err
		}
	}

	srcFile, err := os.Open(s.absPath)
	if err != nil {
		return err
	}
	defer func() { _ = srcFile.Close() }()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer func() { _ = gr.Close() }()
	tr := tar.NewReader(gr)

	return doUnTarGz(tr, absDstPath)
}

func doUnTarGz(tr *tar.Reader, absPath string) error {
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		if hdr.FileInfo().IsDir() {
			_ = os.Mkdir(path.Join(absPath, hdr.Name), 0755)
			continue
		}

		f, err := os.Create(path.Join(absPath, hdr.Name))
		if err != nil {
			continue
		}

		_, err = io.Copy(f, tr)
		if err != nil {
			_ = f.Close()
			return err
		}

		_ = f.Close()
	}

	return nil
}
