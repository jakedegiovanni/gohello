package di

import (
	"net/http"

	"github.com/jakedegiovanni/gohello/internal/app/server"
	"github.com/jakedegiovanni/gohello/internal/app/world"
)

// MakeServer ...
func MakeServer(port int) (*http.Server, error) {
	containers := makeContainers()
	return server.Build(port, containers)
}

func makeContainers() []server.HandlerContainer {
	return []server.HandlerContainer{
		world.NewHandlerContainer(server.NewHandlerContainer),
	}
}
