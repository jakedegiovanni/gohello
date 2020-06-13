package world

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jakedegiovanni/gohello/internal/app/server"
)

const (
	endpoint = "/world/"

	worldGreeting = "Hello, World."
)

var validPath = endpoint + `[a-z]{1,25}/?$`

type response struct {
	Message string `json:"message"`
}

// NewHandlerContainer ...
func NewHandlerContainer(constructor server.HandlerContainerConstructor) server.HandlerContainer {
	return constructor(endpoint, &handler{})
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	re := regexp.MustCompile(validPath)
	var message string
	if path == endpoint {
		message = worldGreeting
	} else if re.MatchString(path) {
		message = getMessageFromPath(path)
	} else {
		http.Error(w, fmt.Sprintf("%s not a valid path. Must be of form %s", path, validPath), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		get(w, message)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func getMessageFromPath(path string) (msg string) {
	msg = strings.ReplaceAll(path, endpoint, "")
	msg = strings.ReplaceAll(msg, "/", "")
	return
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
