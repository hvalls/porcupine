package stream

import "porcupine/storage"

type StreamService struct {
	reader   streamReader
	appender streamAppender
}

func NewStreamService(s storage.StorageService) StreamService {
	return StreamService{
		newStreamReader(s),
		newStreamAppender(s),
	}
}

func (s StreamService) Read(streamId StreamId) (*[]Event, error) {
	return s.reader.read(streamId)
}

func (s StreamService) Append(streamId StreamId, ee []EventWriteModel) error {
	return s.appender.append(ee)
}
