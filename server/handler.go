package server

import (
	"encoding/json"
	"net/http"
	"porcupine/event"
	"porcupine/stream"

	"github.com/gorilla/mux"
)

func handleGetEvents(s event.EventService, w http.ResponseWriter, r *http.Request) {
	streamId, ok := getStreamId(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eventsRead, err := s.Read(stream.StreamId(streamId))
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

func handlePostEvents(s event.EventService, w http.ResponseWriter, r *http.Request) {
	streamId, ok := getStreamId(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var e event.Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.Append(stream.StreamId(streamId), []event.Event{e})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{ \"success\": true }"))
}

func getStreamId(r *http.Request) (string, bool) {
	vars := mux.Vars(r)
	streamId, ok := vars["streamId"]
	return streamId, ok
}
