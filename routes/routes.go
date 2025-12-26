package routes

import (
	"net/http"

	"github.com/radovskyb/not-a-rule-engine/handler"
)

func Setup(h *handler.Handler) *http.ServeMux {
	// Base handler.
	mux := http.NewServeMux()

	mux.Handle("POST /api/ingest", h.Serve(h.Ingest))

	return mux
}
