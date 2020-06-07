package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	methodNotSupportMessage = "%s Operations not support on this endpoint"

	noHandlers = "No handlers supplied"
	noPort     = "No port supplied"
)

// NotImplemented ...
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, methodNotSupportMessage, r.Method)
}

// HandlerContainer ...
type HandlerContainer struct {
	Endpoint string
	Handler  http.Handler
}

// NewBuilder ...
func NewBuilder() Builder {
	return &builder{}
}

// Builder ...
type Builder interface {
	Build(port int, containers []HandlerContainer) (*http.Server, error)
}

type builder struct{}

// NewServer ...
func (b *builder) Build(port int, containers []HandlerContainer) (*http.Server, error) {
	if port == 0 {
		return nil, errors.New(noPort)
	}

	handler, err := newHandler(containers)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}, nil
}

func newHandler(containers []HandlerContainer) (http.Handler, error) {
	if len(containers) == 0 {
		return nil, errors.New(noHandlers)
	}
	mux := http.NewServeMux()
	for _, container := range containers {
		mux.Handle(container.Endpoint, container.Handler)
	}
	return mux, nil
}
