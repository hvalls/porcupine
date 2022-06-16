package buffer

import (
	"encoding/binary"
	"io"
	"strings"
)

func WriteUint32(w io.Writer, data uint32) error {
	arr := make([]byte, 4)
	binary.BigEndian.PutUint32(arr, data)
	_, err := w.Write(arr)
	return err
}

func ReadRangeUint32(data []byte, from int, to int) uint32 {
	v := binary.BigEndian.Uint32(data[from:to])
	return v
}

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

func ReadRangeString(data []byte, from int, to int, trimZeros bool) string {
	rangeData := string(data[from:to])
	if !trimZeros {
		return rangeData
	}
	return strings.Trim(rangeData, string(rune(0)))
}
