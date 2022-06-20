package file

import (
	"os"
)

func OpenWritable(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0600)
}

func OpenReadable(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_RDONLY, 0600)
}
