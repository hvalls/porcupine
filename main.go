package main

import (
	"os"
	"porcupine/server"
	"porcupine/storage"
	"porcupine/stream"
)

func main() {
	storageService := storage.NewStorageService()
	streamService := stream.NewStreamService(storageService)
	server := server.NewServer(streamService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Listen(port)
}
