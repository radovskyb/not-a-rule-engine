package routes

import (
	"net/http"

	"github.com/radovskyb/not-a-rule-engine/handler"
)

func Setup(h *handler.Handler) *http.ServeMux {
	// Base handler.
	mux := http.NewServeMux()

	mux.Handle("/api/ingest", h.Serve(h.Ingest))

	return mux
}

// Oh geez, I can't remember, but I'm starting to think the default router doesn't have fncs for request type like gorilla/mux.
//
// TODO: Create a wrapper to differentiate between POST, GET, etc.
