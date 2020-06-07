package main

import (
	"flag"
	"log"

	"github.com/jakedegiovanni/gohello/internal/app"
)

const defaultPort = 8080

type opts struct {
	port int
}

var parseFlags = func() *opts {
	var port int
	flag.IntVar(&port, "port", defaultPort, "Port to run the server on.")
	flag.Parse()

	return &opts{port: port}
}

func main() {
	opt := parseFlags()

	containers := app.MakeContainers()
	builder := app.ServerBuilder()
	server := app.App{
		Port:       opt.port,
		Containers: containers,
		Builder:    builder,
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
