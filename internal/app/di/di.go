package di

import (
	"net/http"

	"github.com/jakedegiovanni/gohello/internal/app/server"
	"github.com/jakedegiovanni/gohello/internal/app/world"
)

// MakeServer ...
func MakeServer(port int) (*http.Server, error) {
	handlerMap := makeHandlerMap()
	frontController := server.NewFrontController(handlerMap)
	return server.NewServer(port, frontController)
}

func makeHandlerMap() map[string]http.Handler {
	return map[string]http.Handler{
		world.Endpoint: world.Handler(),
	}
}
