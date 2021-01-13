package main

import (
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/subintp/twirp/internal/haberdasherserver"
	"github.com/subintp/twirp/rpc/haberdasher"
	"github.com/subintp/twirp/tracing"
	ottwirp "github.com/twirp-ecosystem/twirp-opentracing"
)

func main() {
	tracing.InitTracer("twirp-server")
	tracer := opentracing.GlobalTracer()
	hooks := ottwirp.NewOpenTracingHooks(tracer)

	service := &haberdasherserver.Server{} // implements Haberdasher interface
	server := ottwirp.WithTraceContext(haberdasher.NewHaberdasherServer(service, hooks), tracer)

	http.ListenAndServe(":8080", server)
}
