package server

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"
)

var errNoPort = errors.New("No port supplied")

// ContainerConstructor ...
type ContainerConstructor func(string, http.Handler) Container

// NewContainer ...
func NewContainer(endpoint string, handler http.Handler) Container {
	return Container{Endpoint: endpoint, Handler: handler}
}

// Container ...
type Container struct {
	Endpoint string
	Handler  http.Handler
}

// NewFrontController ...
func NewFrontController(containers ...Container) *FrontController {
	x := make(map[string]Container)
	for _, e := range containers {
		if _, ok := x[e.Endpoint]; ok {
			panic(fmt.Sprintf("Duplicated top level endpoint found: %s", e.Endpoint))
		} else {
			x[e.Endpoint] = e
		}
	}
	return &FrontController{endpoints: x}
}

// FrontController ...
type FrontController struct {
	endpoints map[string]Container
}

func (f *FrontController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	if container, ok := f.endpoints[head]; ok {
		container.Handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// PathShift ...
type PathShift func(string) (head, tail string)

// ShiftPath https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
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
