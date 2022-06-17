package chunk

import (
	"io"
	"porcupine/buffer"
)

const chunkHeaderSize = 8192 //8kb
const eventCountSize = 1

type chunkHeader struct {
	eventCount uint8
}

func (h chunkHeader) write(w io.Writer) error {
	err := buffer.WriteUint8(w, h.eventCount)
	if err != nil {
		return err
	}
	return nil
}

func read(r io.Reader) (*chunkHeader, error) {
	data, err := buffer.Read(r, chunkHeaderSize)
	if err != nil {
		return nil, err
	}
	eventCount := buffer.ReadRangeUint8(data, 0, eventCountSize)
	return &chunkHeader{eventCount}, nil
}
