package stream

import (
	"porcupine/storage"
)

type streamAppender struct {
	s storage.StorageService
}

func newStreamAppender(s storage.StorageService) streamAppender {
	return streamAppender{s}
}

func (w *streamAppender) append(evs []EventWriteModel) error {
	for _, ev := range evs {
		err := w.s.Append(string(ev.Id), string(ev.StreamId), string(ev.Type), ev.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
