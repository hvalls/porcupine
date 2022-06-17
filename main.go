package main

import (
	"os"
	"porcupine/record"
	"porcupine/server"
	"porcupine/stream"
)

func main() {
	recordService := record.NewRecordService()
	streamService := stream.NewStreamService(recordService)
	server := server.NewServer(streamService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Listen(port)
}
