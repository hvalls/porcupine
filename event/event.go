package event

import "porcupine/stream"

type Event struct {
	Id       EventId         `json:"id"`
	StreamId stream.StreamId `json:"streamId"`
	Type     EventType       `json:"type"`
	Data     EventData       `json:"data"`
}

type EventId string
type EventType string
type EventData []byte

func NewEvent(id EventId, s stream.StreamId, t EventType, p EventData) Event {
	return Event{id, s, t, p}
}
