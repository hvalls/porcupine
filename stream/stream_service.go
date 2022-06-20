package stream

import "porcupine/storage"

type StreamService struct {
	reader   streamReader
	appender streamAppender
	manager  streamManager
}

func NewStreamService(s storage.StorageService) StreamService {
	return StreamService{
		newStreamReader(s),
		newStreamAppender(s),
		newStreamManager(s),
	}
}

func (s StreamService) Create(streamId StreamId) error {
	return s.manager.create(streamId)
}

func (s StreamService) Read(streamId StreamId) (*[]Event, error) {
	return s.reader.read(streamId)
}

func (s StreamService) Append(streamId StreamId, ee []EventWriteModel) error {
	return s.appender.append(ee)
}
