package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"porcupine/stream"
)

func (s Server) handleStreams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.handlePostStream(w, r)
	}
}

func (s Server) handlePostStream(w http.ResponseWriter, r *http.Request) {
	var srm StreamReqModel
	err := json.NewDecoder(r.Body).Decode(&srm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.s.Create(stream.StreamId(srm.Id))
	if err != nil {

		if err.Error() == "StreamAlreadyExists" {
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", srm.Id)))
}
