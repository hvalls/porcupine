package file

import (
	"bufio"
	"io"
	"os"
)

func NewFileWriter(filename string) (*bufio.Writer, *os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
