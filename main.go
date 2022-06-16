package main

import (
	"porcupine/event"
	"porcupine/record"
	"porcupine/server"
)

func main() {
	recordService := record.NewRecordService()
	eventService := event.NewEventService(recordService)

	server.Listen(eventService)
}
