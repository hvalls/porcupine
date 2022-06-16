package record

import (
	"io"
	"porcupine/buffer"
)

const EventIdSize = 32
const StreamIdSize = 32
const EventTypeSize = 32
const EventDataSize = 65536 //64kb
const RecordSize = EventIdSize + StreamIdSize + EventTypeSize + EventDataSize

type Record struct {
	EventId   string
	StreamId  string
	EventType string
	EventData []byte
}

func NewRecord(eventId string, streamId string, eventType string, eventData []byte) Record {
	return Record{
		EventId:   eventId,
		StreamId:  streamId,
		EventType: eventType,
		EventData: eventData,
	}
}

func (r Record) Write(w io.Writer) error {
	err := buffer.Write(w, []byte(r.EventId))
	if err != nil {
		return err
	}

	err = buffer.WriteExpanded(w, []byte(r.StreamId), StreamIdSize)
	if err != nil {
		return err
	}

	err = buffer.WriteExpanded(w, []byte(r.EventType), EventTypeSize)
	if err != nil {
		return err
	}

	err = buffer.WriteExpanded(w, r.EventData, EventDataSize)
	if err != nil {
		return err
	}

	return nil
}

func ReadNext(r io.Reader) (*Record, error) {
	data, err := buffer.Read(r, RecordSize)
	if err != nil {
		return nil, err
	}
	eventId := buffer.ReadRange(data, 0, EventIdSize, false)
	streamId := buffer.ReadRange(data, EventIdSize, EventIdSize+StreamIdSize, true)
	eventType := buffer.ReadRange(data, EventIdSize+StreamIdSize, EventIdSize+StreamIdSize+EventTypeSize, true)
	eventData := buffer.ReadRange(data, EventIdSize+StreamIdSize+EventTypeSize, EventIdSize+StreamIdSize+EventTypeSize+EventDataSize, true)
	return &Record{EventId: eventId, StreamId: streamId, EventType: eventType, EventData: []byte(eventData)}, nil
}
