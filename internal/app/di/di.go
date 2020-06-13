package di

import (
	"net/http"

	"github.com/jakedegiovanni/gohello/internal/app/server"
	"github.com/jakedegiovanni/gohello/internal/app/world"
)

var containerInitialisers = []server.HandlerContainerBuilder{
	world.NewHandlerContainer,
}

// MakeServer ...
func MakeServer(port int) (*http.Server, error) {
	containers := makeContainers(containerInitialisers)
	return server.Build(port, containers)
}

func makeContainers(builders []server.HandlerContainerBuilder) []server.HandlerContainer {
	var containers []server.HandlerContainer
	for _, builder := range builders {
		containers = append(containers, builder(server.NewHandlerContainer))
	}
	return containers
}
