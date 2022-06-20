package file

import (
	"bufio"
	"io"
	"os"
)

func OpenWritable(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0600)
}

func NewFileWriter(filename string) (*bufio.Writer, *os.File, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, nil, err
	}
	w := bufio.NewWriter(f)
	return w, f, nil
}

func NewFileReader(filename string) (io.Reader, *os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	return f, f, nil
}
