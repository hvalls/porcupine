package storage

import (
	"io"
	"porcupine/buffer"
)

const EventIdSize = 36 //UUID
const EventNumberSize = 4
const StreamIdSize = 32
const EventTypeSize = 32
const EventDataSize = 65536 //64kb
const RecordSize = EventIdSize +
	EventNumberSize +
	StreamIdSize +
	EventTypeSize +
	EventDataSize

type Record struct {
	EventId     string
	EventNumber uint32
	StreamId    string
	EventType   string
	EventData   []byte
}

func NewRecord(eventId string, eventNumber uint32, streamId string, eventType string, eventData []byte) Record {
	return Record{
		EventId:     eventId,
		EventNumber: eventNumber,
		StreamId:    streamId,
		EventType:   eventType,
		EventData:   eventData,
	}
}

func (r Record) Write(w io.Writer) error {
	err := buffer.Write(w, []byte(r.EventId))
	if err != nil {
		return err
	}

	err = buffer.WriteUint32(w, r.EventNumber)
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
	eventId := buffer.ReadRangeString(data, 0, EventIdSize, false)
	eventNumber := buffer.ReadRangeUint32(data, EventIdSize, EventIdSize+EventNumberSize)
	streamId := buffer.ReadRangeString(data, EventIdSize+EventNumberSize, EventIdSize+EventNumberSize+StreamIdSize, true)
	eventType := buffer.ReadRangeString(data, EventIdSize+EventNumberSize+StreamIdSize, EventIdSize+EventNumberSize+StreamIdSize+EventTypeSize, true)
	eventData := buffer.ReadRangeString(data, EventIdSize+EventNumberSize+StreamIdSize+EventTypeSize, EventIdSize+EventNumberSize+StreamIdSize+EventTypeSize+EventDataSize, true)
	return &Record{EventId: eventId, EventNumber: eventNumber, StreamId: streamId, EventType: eventType, EventData: []byte(eventData)}, nil
}
