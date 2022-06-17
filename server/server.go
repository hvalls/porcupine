package server

import (
	"fmt"
	"net/http"
	"porcupine/event"

	"github.com/gorilla/mux"
)

type Server struct {
	s event.EventService
}

func NewServer(s event.EventService) Server {
	return Server{s}
}

func (s Server) Listen(port string) {
	r := mux.NewRouter()
	r.Path("/streams/{streamId}/events").Methods(http.MethodGet, http.MethodPost).HandlerFunc(s.handleEvents)

	fmt.Printf("server listening at port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
