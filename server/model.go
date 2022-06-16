package server

import (
	"porcupine/event"
	"porcupine/stream"
)

type EventReqModel struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (erm EventReqModel) Event(streamId string) event.Event {
	return event.NewEvent(
		event.EventId(erm.Id),
		stream.StreamId(streamId),
		event.EventType(erm.Type),
		event.EventData(erm.Data),
	)
}
