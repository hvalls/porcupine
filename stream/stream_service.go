package stream

import (
	"porcupine/record"
)

type StreamService struct {
	reader   streamReader
	appender streamAppender
}

func NewStreamService(s record.RecordService) StreamService {
	return StreamService{
		newStreamReader(s),
		newStreamAppender(s),
	}
}

func (s StreamService) Read(streamId StreamId) (*[]Event, error) {
	return s.reader.Read(streamId)
}

func (s StreamService) Append(streamId StreamId, ee []EventWriteModel) error {
	return s.appender.Append(ee)
}
