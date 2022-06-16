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

func (w *EventAppender) Append(evs []EventWriteModel) error {
	for _, ev := range evs {
		err := w.s.Append(string(ev.Id), string(ev.StreamId), string(ev.Type), ev.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
