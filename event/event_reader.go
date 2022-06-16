package event

import (
	"porcupine/record"
	"porcupine/stream"
)

type EventReader struct {
	s record.RecordService
}

func NewEventReader(s record.RecordService) EventReader {
	return EventReader{s}
}

func (r *EventReader) Read(streamId stream.StreamId) (*[]Event, error) {
	records, err := r.s.Read(string(streamId))
	if err != nil {
		return nil, err
	}
	events := make([]Event, 0)
	for _, record := range records {
		event := New(
			EventId(record.EventId),
			EventNumber(record.EventNumber),
			stream.StreamId(record.StreamId),
			EventType(record.EventType),
			EventData(record.EventData),
		)
		events = append(events, event)
	}
	return &events, nil
}
