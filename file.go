/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-01-02 15:59
**/

package utils

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type fi int

const File fi = iota

type fileInfo struct {
	reader io.Reader
	err    error
}

func (fi fi) IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (fi fi) ReadFromBytes(bts []byte) fileInfo {
	return fileInfo{err: nil, reader: bytes.NewReader(bts)}
}

func (fi fi) ReadFromString(str string) fileInfo {
	return fileInfo{err: nil, reader: strings.NewReader(str)}
}

func (fi fi) ReadFromReader(r io.Reader) fileInfo {
	return fileInfo{err: nil, reader: r}
}

func (i fileInfo) Error() error {
	return i.err
}

func (i fileInfo) Bytes() []byte {
	if i.err != nil {
		return nil
	}
	var bts, err = io.ReadAll(i.reader)
	if err != nil {
		return nil
	}
	return bts
}

func (i fileInfo) String() string {
	if i.err != nil {
		return ""
	}
	var bts, err = io.ReadAll(i.reader)
	if err != nil {
		return ""
	}
	return string(bts)
}

func (i fileInfo) WriteToPath(path string) error {
	if i.err != nil {
		return i.err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	f, err := os.Create(absPath)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = io.Copy(f, i.reader)
	if err != nil {
		return err
	}

	return nil
}

func (i fileInfo) WriteToReader(w io.Writer) error {
	if i.err != nil {
		return i.err
	}

	_, err := io.Copy(w, i.reader)
	if err != nil {
		return err
	}
	return nil
}
