/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2020-01-02 15:59
**/

package file

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Info struct {
	reader io.Reader
	err    error
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func FromBytes(bts []byte) Info {
	return Info{err: nil, reader: bytes.NewReader(bts)}
}

func FromString(str string) Info {
	return Info{err: nil, reader: strings.NewReader(str)}
}

func FromReader(r io.Reader) Info {
	return Info{err: nil, reader: r}
}

func (i Info) Error() error {
	return i.err
}

func (i Info) Bytes() []byte {
	if i.err != nil {
		return nil
	}
	var bts, err = io.ReadAll(i.reader)
	if err != nil {
		return nil
	}
	return bts
}

func (i Info) String() string {
	if i.err != nil {
		return ""
	}
	var bts, err = io.ReadAll(i.reader)
	if err != nil {
		return ""
	}
	return string(bts)
}

func (i Info) WriteToPath(path string) error {
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

func (i Info) WriteToReader(w io.Writer) error {
	if i.err != nil {
		return i.err
	}

	_, err := io.Copy(w, i.reader)
	if err != nil {
		return err
	}
	return nil
}
