package stream

import "porcupine/storage"

type streamReader struct {
	s storage.StorageService
}

func newStreamReader(s storage.StorageService) streamReader {
	return streamReader{s}
}

func (r *streamReader) Read(streamId StreamId) (*[]Event, error) {
	records, err := r.s.Read(string(streamId))
	if err != nil {
		return nil, err
	}
	events := make([]Event, 0)
	for _, record := range records {
		event := New(
			EventId(record.EventId),
			EventNumber(record.EventNumber),
			StreamId(record.StreamId),
			EventType(record.EventType),
			EventData(record.EventData),
		)
		events = append(events, event)
	}
	return &events, nil
}
