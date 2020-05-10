package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jakedegiovanni/gohello/pkg/world"
)

type genericResponse struct {
	Message string `json:"message"`
}

// Start ...
func Start() {
	http.HandleFunc("/helloworld", worldHandler)
	fmt.Printf("Starting Server on port 8080\n")
	http.ListenAndServe(":8080", nil)
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
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%s Operations not support", r.Method)
	}
}
