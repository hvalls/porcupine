package stream

import (
	"porcupine/record"
)

type streamAppender struct {
	s record.RecordService
}

func newStreamAppender(s record.RecordService) streamAppender {
	return streamAppender{s}
}

func (w *streamAppender) Append(evs []EventWriteModel) error {
	for _, ev := range evs {
		err := w.s.Append(string(ev.Id), string(ev.StreamId), string(ev.Type), ev.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
