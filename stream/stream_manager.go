package stream

import (
	"porcupine/storage"
)

type streamManager struct {
	s storage.StorageService
}

func newStreamManager(s storage.StorageService) streamManager {
	return streamManager{s}
}

func (w *streamManager) create(streamId StreamId) error {
	return w.s.CreateStream(string(streamId))
}
