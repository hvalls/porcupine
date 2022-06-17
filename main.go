package main

import (
	"os"
	"porcupine/event"
	"porcupine/record"
	"porcupine/server"
)

func main() {
	recordService := record.NewRecordService()
	eventService := event.NewEventService(recordService)
	server := server.NewServer(eventService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Listen(port)
}
