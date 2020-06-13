package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	errNoHandler = errors.New("No handlers supplied")
	errNoPort    = errors.New("No port supplied")
)

// HandlerContainerBuilder ...
type HandlerContainerBuilder func(HandlerContainerConstructor) HandlerContainer

// HandlerContainerConstructor ...
type HandlerContainerConstructor func(endpoint string, handler http.Handler) HandlerContainer

// NewHandlerContainer ...
func NewHandlerContainer(endpoint string, handler http.Handler) HandlerContainer {
	return HandlerContainer{Endpoint: endpoint, Handler: handler}
}

// HandlerContainer ...
type HandlerContainer struct {
	Endpoint string
	Handler  http.Handler
}

// Build ...
func Build(port int, containers []HandlerContainer) (*http.Server, error) {
	if port == 0 {
		return nil, errNoPort
	}
	if len(containers) == 0 {
		return nil, errNoHandler
	}

	handler := newHandler(containers)

	return &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}, nil
}

func newHandler(containers []HandlerContainer) http.Handler {
	mux := http.NewServeMux()
	for _, container := range containers {
		mux.Handle(container.Endpoint, container.Handler)
	}
	return mux
}
