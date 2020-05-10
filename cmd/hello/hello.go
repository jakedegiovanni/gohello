package main

import (
	"flag"

	"github.com/jakedegiovanni/gohello/pkg/server"
)

func main() {
	var portPtr *int = flag.Int("port", 8080, "Port to run the server on.")
	flag.Parse()

	server.Start(*portPtr)
}
