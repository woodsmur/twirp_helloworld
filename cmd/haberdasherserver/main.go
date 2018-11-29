package main

import (
	"net/http"

	"github.com/woodsmur/twirp_helloworld/internal/haberdasherserver"
	"github.com/woodsmur/twirp_helloworld/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}
