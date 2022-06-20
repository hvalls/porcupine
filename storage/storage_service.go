package storage

import (
	"errors"
	"fmt"
	"io"
	"porcupine/chunk"
	"porcupine/file"
	"porcupine/record"
	"strings"
)

const streamsFile = "streams"

type StorageService struct {
}

func NewStorageService() StorageService {
	return StorageService{}
}

func (s StorageService) CreateStream(streamId string) error {
	streams, err := readStreams()
	if err != nil {
		return err
	}
	for _, s := range streams {
		if s == streamId {
			return errors.New("StreamAlreadyExists")
		}
	}
	err = writeNewStream(streamId)
	if err != nil {
		return err
	}
	return chunk.CreateChunk(streamId)
}

func (s StorageService) StoreRecord(
	eventId string,
	streamId string,
	eventType string,
	eventData []byte,
) error {
	return chunk.GetChunk(streamId).WriteRecord(eventId, eventType, eventData)
}

func (s StorageService) GetRecords(streamId string) ([]record.Record, error) {
	return chunk.GetChunk(streamId).ReadRecords()
}

func readStreams() ([]string, error) {
	f, err := file.OpenReadable(streamsFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	lines := string(data)

	return strings.Split(lines, "\n"), nil
}

func writeNewStream(streamId string) error {
	f, err := file.OpenWritableAppend(streamsFile)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(fmt.Sprintf("%s\n", streamId)))
	return err
}
