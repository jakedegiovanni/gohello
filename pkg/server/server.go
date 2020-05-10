package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jakedegiovanni/gohello/pkg/world"
)

type genericResponse struct {
	Message string `json:"message"`
}

// Start ...
func Start(port int) {
	http.HandleFunc("/helloworld", worldHandler)

	fmt.Printf("Starting Server on port %d\n", port)

	addr := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func worldHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		resp := world.Hello()
		jsonResp := genericResponse{
			Message: resp,
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)

		json.NewEncoder(rw).Encode(jsonResp)
	} else {
		rw.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(rw, "%s Operations not support on this endpoint", r.Method)
	}
}
