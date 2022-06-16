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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eventsRead, err := srv.Read(stream.StreamId(streamId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(eventsRead)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("content-type", "application/json")
	w.Write(j)
}
