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
	"io/ioutil"
	"os"
	"path/filepath"
)

type fi int

const File fi = iota

type fileInfo struct {
	bytes []byte
	err   error
}

func (fi fi) IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (fi fi) ReadFromBytes(bts []byte) fileInfo {
	return fileInfo{err: nil, bytes: bts}
}

func (fi fi) ReadFromString(str string) fileInfo {
	return fileInfo{err: nil, bytes: []byte(str)}
}

func (fi fi) ReadFromReader(r io.Reader) fileInfo {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fileInfo{err: err, bytes: nil}
	}
	return fileInfo{err: nil, bytes: b}
}

func (fi fi) ReadFromPath(path string) fileInfo {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fileInfo{err: err, bytes: nil}
	}
	b, err := ioutil.ReadFile(absPath)
	if err != nil {
		return fileInfo{err: err, bytes: nil}
	}
	return fileInfo{err: nil, bytes: b}
}

func (i fileInfo) LastError() error {
	return i.err
}

func (i fileInfo) Bytes() []byte {
	return i.bytes
}

func (i fileInfo) Append(bts []byte) fileInfo {
	i.bytes = append(i.bytes, bts...)
	return i
}

func (i fileInfo) Slice(start int, end int) fileInfo {
	i.bytes = i.bytes[start:end]
	return i
}

func (i fileInfo) String() string {
	return string(i.bytes)
}

func (i fileInfo) WriteToPath(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	f, err := os.Create(absPath)
	defer func() { _ = f.Close() }()
	if err != nil {
		return err
	}

	_, err = io.Copy(f, bytes.NewReader(i.bytes))
	if err != nil {
		return err
	}

	return nil
}

func (i fileInfo) WriteToReader(w io.Writer) error {
	_, err := io.Copy(w, bytes.NewReader(i.bytes))
	if err != nil {
		return err
	}
	return nil
}

func (i fileInfo) WriteToBytes(bts []byte) error {
	if len(bts) <= len(i.bytes) {
		copy(bts, i.bytes)
		return nil
	}
	bts = bts[0 : len(i.bytes)-1]
	copy(bts, i.bytes)
	return nil
}
