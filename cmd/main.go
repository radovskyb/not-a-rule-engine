package main

import (
	"log"
	"net/http"

	"github.com/radovskyb/not-a-rule-engine/handler"
)

func main() {
	// Setup.
	//
	// Handler
	h := handler.New()

	// For setup sake, just creating a fake route in main, but going to move all of this out into it's own package soonish.
	// This isn't a regular http.HandlerFunc, but my custom handler.HandlerFunc type
	hf := func(w http.ResponseWriter, r *http.Request) error {
		// Do stuff.
		log.Println("Setup test")
		return nil
	}

	mux := http.NewServeMux()
	mux.Handle("/api", h.Serve(hf))

	http.ListenAndServe(":9000", mux)
}
