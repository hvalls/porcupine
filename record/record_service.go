package record

type RecordService struct {
}

func NewRecordService() RecordService {
	return RecordService{}
}

func (s RecordService) Append(record Record) error {
	return GetChunk(record.StreamId).Write(record)
}

func (s RecordService) Read(streamId string) ([]Record, error) {
	return GetChunk(streamId).Read()
}
