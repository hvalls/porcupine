package file

import (
	"os"
)

func OpenWritableAppend(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
}

func OpenWritable(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0600)
}

func OpenReadable(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0600)
}
