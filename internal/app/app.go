package app

import (
	"fmt"

	"github.com/jakedegiovanni/gohello/internal/app/server"
	"github.com/jakedegiovanni/gohello/internal/app/world"
)

// MakeContainers ...
func MakeContainers() []server.HandlerContainer {
	return []server.HandlerContainer{world.NewHandlerContainer()}
}

// ServerBuilder ...
func ServerBuilder() server.Builder {
	return server.NewBuilder()
}

// App ...
type App struct {
	Port       int
	Containers []server.HandlerContainer
	Builder    server.Builder
}

// Start ...
func (a *App) Start() error {
	server, err := a.Builder.Build(a.Port, a.Containers)
	if err != nil {
		return err
	}
	fmt.Printf("Starting Server on address %s\n", server.Addr)
	return server.ListenAndServe()
}
