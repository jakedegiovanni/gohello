package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jakedegiovanni/gohello/internal/app/di"
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

	server, err := di.MakeServer(opt.port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting Server on address %s\n", server.Addr)
	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
