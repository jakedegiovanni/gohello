package world

import (
	"encoding/json"
	"net/http"

	"github.com/jakedegiovanni/gohello/internal/app/server/urlutil"
)

const (
	// Endpoint ...
	Endpoint = "world"

	worldGreeting = "Hello, World."
)

// Handler ...
func Handler() http.Handler {
	return &handler{}
}

type response struct {
	Message string `json:"message"`
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// valid paths: "/" , "/:message[/]+"
	// invalid paths: "/:message/[a-Z]*[/]+", "/:message/[a-Z]*"
	head, tail := urlutil.ShiftPath(r.URL.Path)
	if tail != "/" {
		http.NotFound(w, r)
		return
	}

	var message string
	if head == "" {
		message = worldGreeting
	} else {
		message = head
	}

	switch r.Method {
	case http.MethodGet:
		get(w, message)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func get(w http.ResponseWriter, message string) {
	resp := response{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
