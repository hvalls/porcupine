package server

import (
	"fmt"
	"net/http"
	"porcupine/event"

	"github.com/gorilla/mux"
)

func Listen(s event.EventService) {
	r := mux.NewRouter()

	r.HandleFunc("/streams/{streamId}/events", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetEvents(s, w, r)
		default:
		}
	})

	fmt.Println("server listening at port 8080")
	http.ListenAndServe(":8080", r)
}
