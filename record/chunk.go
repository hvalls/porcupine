package record

import (
	"fmt"
	"io"
	"porcupine/file"
)

type Chunk struct {
	filename string
}

func GetChunk(streamId string) Chunk {
	return Chunk{resolveFileName(streamId)}
}

func (c Chunk) Write(r Record) error {
	w, f, err := file.NewFileWriter(c.filename)
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

func (c Chunk) Read() ([]Record, error) {
	r, f, err := file.NewFileReader(c.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var records []Record
	for {
		record, err := ReadNext(r)
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
