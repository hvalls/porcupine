package buffer

import (
	"io"
	"strings"
)

func Write(w io.Writer, data []byte) error {
	_, err := w.Write([]byte(data))
	return err
}

func WriteExpanded(w io.Writer, data []byte, size int) error {
	bb := make([]byte, size)

	from := len(bb) - 1
	lenDiff := len(bb) - len(data)
	for i := from; i >= lenDiff; i-- {
		bb[i] = data[i-lenDiff]
	}

	_, err := w.Write(bb)
	if err != nil {
		return err
	}

	return nil
}

func Read(r io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	_, err := r.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadRange(data []byte, from int, to int, trimZeros bool) string {
	rangeData := string(data[from:to])
	if !trimZeros {
		return rangeData
	}
	return strings.Trim(rangeData, string(rune(0)))
}
