package record

import (
	"io"
	"porcupine/buffer"
)

const eventIdSize = 36 //UUID
const eventNumberSize = 4
const streamIdSize = 32
const eventTypeSize = 32
const eventDataSize = 65536 //64kb
const recordSize = eventIdSize +
	eventNumberSize +
	streamIdSize +
	eventTypeSize +
	eventDataSize

type Record struct {
	EventId     string
	EventNumber uint32
	StreamId    string
	EventType   string
	EventData   []byte
}

func New(eventId string, eventNumber uint32, streamId string, eventType string, eventData []byte) Record {
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

	err = buffer.WriteExpanded(w, []byte(r.StreamId), streamIdSize)
	if err != nil {
		return err
	}

	err = buffer.WriteExpanded(w, []byte(r.EventType), eventTypeSize)
	if err != nil {
		return err
	}

	err = buffer.WriteExpanded(w, r.EventData, eventDataSize)
	if err != nil {
		return err
	}

	return nil
}

func Read(r io.Reader) (*Record, error) {
	data, err := buffer.Read(r, recordSize)
	if err != nil {
		return nil, err
	}
	eventId := buffer.ReadRangeString(data, 0, eventIdSize, false)
	eventNumber := buffer.ReadRangeUint32(data, eventIdSize, eventIdSize+eventNumberSize)
	streamId := buffer.ReadRangeString(data, eventIdSize+eventNumberSize, eventIdSize+eventNumberSize+streamIdSize, true)
	eventType := buffer.ReadRangeString(data, eventIdSize+eventNumberSize+streamIdSize, eventIdSize+eventNumberSize+streamIdSize+eventTypeSize, true)
	eventData := buffer.ReadRangeString(data, eventIdSize+eventNumberSize+streamIdSize+eventTypeSize, eventIdSize+eventNumberSize+streamIdSize+eventTypeSize+eventDataSize, true)
	return &Record{EventId: eventId, EventNumber: eventNumber, StreamId: streamId, EventType: eventType, EventData: []byte(eventData)}, nil
}
