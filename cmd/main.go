package main

import (
	"net/http"
	"os"

	"github.com/radovskyb/not-a-rule-engine/dispatcher"
	"github.com/radovskyb/not-a-rule-engine/handler"
	"github.com/radovskyb/not-a-rule-engine/routes"
	"github.com/radovskyb/not-a-rule-engine/services"
	"github.com/radovskyb/not-a-rule-engine/services/cache"
	"github.com/radovskyb/not-a-rule-engine/services/log"
)

func main() {
	// Setup.

	// Dispatcher
	d := dispatcher.New()

	// Service store
	ss := services.NewStore()

	// Probably this would eventually be imported based on config files.
	cs := cache.New()
	ss.Add(services.CacheServiceID, cs)

	ls := log.New(os.Stdout)
	ss.Add(services.LogServiceID, ls)

	// Handler
	h := handler.New(d, ss)

	// Setup routes.
	mux := routes.Setup(h)

	http.ListenAndServe(":9000", mux)
}
