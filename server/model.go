package server

import (
	"porcupine/event"
	"porcupine/stream"

	"github.com/google/uuid"
)

type EventReqModel struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (erm EventReqModel) WriteModel(streamId string) event.EventWriteModel {
	return event.NewWriteModel(
		event.EventId(uuid.New().String()),
		stream.StreamId(streamId),
		event.EventType(erm.Type),
		event.EventData(erm.Data),
	)
}
