package chunk

import (
	"errors"
	"fmt"
	"io"
	"os"
	"porcupine/file"
	"porcupine/record"
)

type Chunk interface {
	WriteRecord(eventId string, eventType string, eventData []byte) error
	ReadRecords() ([]record.Record, error)
}

type fileChunk struct {
	streamId string
	fileName string
}

func CreateChunk(streamId string) error {
	f, err := file.OpenWritable(resolveFileName(streamId))
	if err != nil {
		return err
	}
	defer f.Close()
	return newChunkHeader(0).write(f)
}

func GetChunk(streamId string) Chunk {
	filename := resolveFileName(streamId)
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return fileChunk{streamId, resolveFileName(streamId)}
}

func (c fileChunk) WriteRecord(eventId string, eventType string, eventData []byte) error {
	f, err := file.OpenWritable(c.fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	h, err := readHeader(f)
	if err != nil {
		return err
	}
	h.incrementEventCount()
	record := record.New(eventId, uint32(h.eventCount), c.streamId, eventType, eventData)
	f.Seek(0, io.SeekEnd)
	err = record.Write(f)
	if err != nil {
		return err
	}
	f.Seek(0, io.SeekStart)
	return h.write(f)
}

func (c fileChunk) ReadRecords() ([]record.Record, error) {
	f, err := file.OpenReadable(c.fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	f.Seek(chunkHeaderSize, io.SeekStart)

	var records []record.Record
	for {
		record, err := record.Read(f)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		records = append(records, *record)
	}

	return records, nil
}

func resolveFileName(streamId string) string {
	return fmt.Sprintf("%s.chunk.1", streamId)
}
