package main

import (
	"net/http"

	"github.com/radovskyb/not-a-rule-engine/dispatcher"
	"github.com/radovskyb/not-a-rule-engine/handler"
	"github.com/radovskyb/not-a-rule-engine/routes"
	"github.com/radovskyb/not-a-rule-engine/services"
	"github.com/radovskyb/not-a-rule-engine/services/cache"
)

func main() {
	// Setup.

	// Dispatcher
	d := dispatcher.New()

	// Service store
	ss := services.NewStore()

	// Example cache service, id = 123
	cs := cache.New()
	ss.Add(123, cs)

	// Handler
	h := handler.New(d, ss)

	// Setup routes.
	mux := routes.Setup(h)

	http.ListenAndServe(":9000", mux)
}
