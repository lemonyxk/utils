/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2020-01-02 16:08
**/

package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Dir struct {
	path string
	err  error
}

func New(path string) Dir {
	var absPath, err = filepath.Abs(path)
	return Dir{path: absPath, err: err}
}

func (d Dir) RemoveAll() error {
	return os.RemoveAll(d.path)
}

func (d Dir) CreateAll(perm os.FileMode) error {
	return os.MkdirAll(d.path, perm)
}

func (d Dir) Create(perm os.FileMode) error {
	return os.Mkdir(d.path, perm)
}

func (d Dir) IsExist() bool {
	_, err := os.Stat(d.path)
	return !os.IsNotExist(err)
}

func (d Dir) Error() error {
	return d.err
}

func (d Dir) ReadAll() []Info {
	var res []Info

	var fn func(path string, res *[]Info)

	fn = func(path string, res *[]Info) {

		files, err := ioutil.ReadDir(path)
		if err != nil {
			*res = append(*res, Info{path, nil, err})
			return
		}

		for i := 0; i < len(files); i++ {
			var fullPath = filepath.Join(path, files[i].Name())
			if files[i].IsDir() {
				fn(fullPath, res)
			}
			*res = append(*res, Info{fullPath, files[i], nil})
		}
	}

	fn(d.path, &res)

	return res
}

func (d Dir) Walk() chan Info {
	var ch = make(chan Info)
	go func() {
		_ = filepath.Walk(d.path, func(path string, info os.FileInfo, err error) error {
			ch <- Info{path, info, err}
			return err
		})
		close(ch)
	}()
	return ch
}

type Info struct {
	path string
	info os.FileInfo
	err  error
}

func (f *Info) Error() error {
	return f.err
}

func (f *Info) Info() os.FileInfo {
	return f.info
}

func (f *Info) Path() string {
	return f.path
}
