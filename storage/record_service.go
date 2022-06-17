package storage

type StorageService struct {
}

func NewStorageService() StorageService {
	return StorageService{}
}

func (s StorageService) Append(
	eventId string,
	streamId string,
	eventType string,
	eventData []byte,
) error {
	chunk := GetChunk(streamId)
	eventNumber := chunk.NextEventNumber()
	r := NewRecord(eventId, eventNumber, streamId, eventType, eventData)
	return chunk.Write(r)
}

func (s StorageService) Read(streamId string) ([]Record, error) {
	return GetChunk(streamId).Read()
}
