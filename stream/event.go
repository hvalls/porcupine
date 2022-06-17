package stream

type Event struct {
	Id       EventId     `json:"id"`
	Number   EventNumber `json:"number"`
	StreamId StreamId    `json:"streamId"`
	Type     EventType   `json:"type"`
	Data     EventData   `json:"data"`
}

type EventId string
type EventNumber int32
type EventType string
type EventData []byte

func New(id EventId, n EventNumber, s StreamId, t EventType, p EventData) Event {
	return Event{id, n, s, t, p}
}

type EventWriteModel struct {
	Id       EventId   `json:"id"`
	StreamId StreamId  `json:"streamId"`
	Type     EventType `json:"type"`
	Data     EventData `json:"data"`
}

func NewWriteModel(id EventId, s StreamId, t EventType, p EventData) EventWriteModel {
	return EventWriteModel{id, s, t, p}
}
