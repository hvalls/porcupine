package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"porcupine/event"
	"porcupine/stream"

	"github.com/gorilla/mux"
)

func (s Server) handleEvents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetEvents(w, r)
	case http.MethodPost:
		s.handlePostEvents(w, r)
	}
}

func (s Server) handleGetEvents(w http.ResponseWriter, r *http.Request) {
	streamId, ok := getStreamId(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eventsRead, err := s.s.Read(stream.StreamId(streamId))
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

func (s Server) handlePostEvents(w http.ResponseWriter, r *http.Request) {
	streamId, ok := getStreamId(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var erm EventReqModel
	err := json.NewDecoder(r.Body).Decode(&erm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e := erm.WriteModel(streamId)
	err = s.s.Append(stream.StreamId(streamId), []event.EventWriteModel{e})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", e.Id)))
}

func getStreamId(r *http.Request) (string, bool) {
	vars := mux.Vars(r)
	streamId, ok := vars["streamId"]
	return streamId, ok
}
