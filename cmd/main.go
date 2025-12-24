package main

import (
	"net/http"

	"github.com/radovskyb/not-a-rule-engine/handler"
	"github.com/radovskyb/not-a-rule-engine/routes"
)

func main() {
	// Setup.
	//
	// Handler
	h := handler.New()

	// Setup routes.
	mux := routes.Setup(h)

	http.ListenAndServe(":9000", mux)
}
