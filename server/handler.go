package server

import (
	"encoding/json"
	"net/http"
	"porcupine/event"
	"porcupine/stream"

	"github.com/gorilla/mux"
)

func handleGetEvents(srv event.EventService, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	streamId, ok := vars["streamId"]
	if !ok {
		panic("invalid streamId param")
	}
	eventsRead, err := srv.Read(stream.StreamId(streamId))
	if err != nil {
		panic(err)
	}
	j, err := json.Marshal(eventsRead)
	if err != nil {
		panic(err)
	}
	w.Header().Add("content-type", "application/json")
	w.Write(j)
}
