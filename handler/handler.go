package handler

import (
	"encoding/json"
	"net/http"

	"github.com/radovskyb/not-a-rule-engine/dispatcher"
	"github.com/radovskyb/not-a-rule-engine/services"
)

type Handler struct {
	d  dispatcher.Client
	ss services.Store
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func New(d dispatcher.Client, ss services.Store) *Handler {
	return &Handler{
		d:  d,
		ss: ss,
	}
}

// Serve ...
//
// Personal preference, but even when not sticking only to stdlib, I like wrapping my handlerfuncs like this for centralized error handling.
func (h *Handler) Serve(hf HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := hf(w, r); err != nil {
			// For now lets just send err, but going to replace this with a custom error type since this is no good for json as is.
			json.NewEncoder(w).Encode(err)
		}
	})
}
