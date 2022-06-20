package storage

import (
	"porcupine/chunk"
	"porcupine/record"
)

type StorageService struct {
}

func NewStorageService() StorageService {
	return StorageService{}
}

func (s StorageService) CreateStream(streamId string) error {
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
