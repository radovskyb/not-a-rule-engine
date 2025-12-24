package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// These would be moved to a shared lookup.
//
// TODO: Move hardcoded val from main.go and add in a config file.
const (
	CacheService = 123
)

type Event struct {
	Type    int `json:"type"`
	Payload any `json:"payload"`
}

// Ingest sounds like a cool name for whatever this is going to end up being. Definitely not for a rule engine tho :(
func (h *Handler) Ingest(w http.ResponseWriter, r *http.Request) error {
	log.Println("nom nom")

	var ev Event
	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		return err
	}

	// NOTE:
	// May want to fan-out to multiple services, so will probably change to accept []Event instead of Event and
	// iterate / pass to worker pool here or something.

	service, err := h.ss.Fetch(ev.Type)
	if err != nil {
		return err
	}

	// Dispatch will handle converting into the correct payload format for the service.
	resp, err := h.d.Dispatch(service, ev.Payload)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(resp)
}
