package world

import (
	"encoding/json"
	"net/http"

	"github.com/jakedegiovanni/gohello/internal/app/server"
)

const (
	endpoint = "/helloworld"

	worldGreeting = "Hello, World."
)

// NewHandlerContainer ...
func NewHandlerContainer() server.HandlerContainer {
	return server.HandlerContainer{
		Endpoint: endpoint,
		Handler:  &handler{},
	}
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		resp := worldGreeting
		jsonResp := server.GenericResponse{
			Message: resp,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		server.NotImplemented(w, r)
	}
}
