package main

import (
	"net/http"

	"github.com/woodsmur/twirp_helloworld/internal/helloworldserver"
	"github.com/woodsmur/twirp_helloworld/rpc/helloworld"
)

func main() {
	server := &helloworldserver.Server{} // implements helloworldserver interface
	twirpHandler := helloworld.NewAwesomeTwirpServiceServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}
