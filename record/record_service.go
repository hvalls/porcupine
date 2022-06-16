package record

type RecordService struct {
}

func NewRecordService() RecordService {
	return RecordService{}
}

func (s RecordService) Append(
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

func (s RecordService) Read(streamId string) ([]Record, error) {
	return GetChunk(streamId).Read()
}
