package main

import (
	"flag"

	"github.com/jakedegiovanni/gohello/pkg/server"
)

var serverStart = server.Start

var defaultPort = 8080

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
	serverStart(opt.port)
}
