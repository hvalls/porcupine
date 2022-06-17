package server

import (
	"porcupine/stream"

	"github.com/google/uuid"
)

type EventReqModel struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (erm EventReqModel) WriteModel(streamId string) stream.EventWriteModel {
	return stream.NewWriteModel(
		stream.EventId(uuid.New().String()),
		stream.StreamId(streamId),
		stream.EventType(erm.Type),
		stream.EventData(erm.Data),
	)
}
