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

func (erm EventReqModel) Event(streamId string) event.Event {
	return event.NewEvent(
		event.EventId(uuid.New().String()),
		stream.StreamId(streamId),
		event.EventType(erm.Type),
		event.EventData(erm.Data),
	)
}
