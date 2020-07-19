package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jakedegiovanni/gohello/internal/app/server/urlutil"
)

var errNoPort = errors.New("No port supplied")

// NewFrontController ...
func NewFrontController(x map[string]http.Handler) *FrontController {
	return &FrontController{handlerMap: x}
}

// FrontController ...
type FrontController struct {
	handlerMap map[string]http.Handler
}

func (f *FrontController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = urlutil.ShiftPath(r.URL.Path)
	if handler, ok := f.handlerMap[head]; ok {
		handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// NewServer ...
func NewServer(port int, frontController *FrontController) (*http.Server, error) {
	if port == 0 {
		return nil, errNoPort
	}

	return &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        frontController,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}, nil
}
