package event

import (
	"porcupine/record"
	"porcupine/stream"
)

type EventService struct {
	reader   EventReader
	appender EventAppender
}

func NewEventService(s record.RecordService) EventService {
	return EventService{
		NewEventReader(s),
		NewEventAppender(s),
	}
}

func (s EventService) Read(streamId stream.StreamId) (*[]Event, error) {
	return s.reader.Read(streamId)
}

func (s EventService) Append(ee []Event) error {
	return s.appender.Append(ee)
}
