package server

import (
	"fmt"
	"net/http"
	"porcupine/stream"

	"github.com/gorilla/mux"
)

type Server struct {
	s stream.StreamService
}

func NewServer(s stream.StreamService) Server {
	return Server{s}
}

func (s Server) Listen(port string) {
	r := mux.NewRouter()
	r.Path("/streams").Methods(http.MethodPost).HandlerFunc(s.handleStreams)
	r.Path("/streams/{streamId}/events").Methods(http.MethodGet, http.MethodPost).HandlerFunc(s.handleEvents)

	fmt.Printf("server listening at port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
