package event

import (
	"porcupine/record"
)

type EventAppender struct {
	s record.RecordService
}

func NewEventAppender(s record.RecordService) EventAppender {
	return EventAppender{s}
}

func (w *EventAppender) Append(evs []Event) error {
	for _, ev := range evs {
		r := record.NewRecord(string(ev.Id), string(ev.StreamId), string(ev.Type), ev.Data)
		err := w.s.Append(r)
		if err != nil {
			return err
		}
	}
	return nil
}
