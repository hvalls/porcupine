package chunk

import (
	"fmt"
	"io"
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

func GetChunk(streamId string) Chunk {
	return fileChunk{streamId, resolveFileName(streamId)}
}

func (c fileChunk) WriteRecord(eventId string, eventType string, eventData []byte) error {

	//TODO implement this
	nextEventNumber := 1

	r := record.New(eventId, uint32(nextEventNumber), c.streamId, eventType, eventData)

	w, f, err := file.NewFileWriter(c.fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	err = r.Write(w)
	if err != nil {
		return err
	}
	return w.Flush()
}

func (c fileChunk) ReadRecords() ([]record.Record, error) {
	r, f, err := file.NewFileReader(c.fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var records []record.Record
	for {
		record, err := record.Read(r)
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
	return fmt.Sprintf("%s.stream.porcupine", streamId)
}
