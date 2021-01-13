package main

import (
	"net/http"

	"github.com/subintp/twirp/internal/haberdasherserver"
	"github.com/subintp/twirp/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
