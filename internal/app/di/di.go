package di

import (
	"net/http"

	"github.com/jakedegiovanni/gohello/internal/app/world"

	"github.com/jakedegiovanni/gohello/internal/app/server"
)

// MakeServer ...
func MakeServer(port int) (*http.Server, error) {
	containers := makeContainers()
	frontController := server.NewFrontController(containers...)
	return server.NewServer(port, frontController)
}

func makeContainers() []server.Container {
	return []server.Container{
		world.NewContainer(server.NewContainer, server.ShiftPath),
	}
}
