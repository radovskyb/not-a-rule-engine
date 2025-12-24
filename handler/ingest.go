package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// May as well start with enums for type. Will replace this, but just for starting sake.

type EventType int

const (
	ServiceOneType EventType = iota
	ServiceTwoType
	ServiceThreeType
)

type Event struct {
	Type int `json:"type"`
}

// Ingest sounds like a cool name for whatever this is going to end up being. Definitely not for a rule engine tho :(
func (h *Handler) Ingest(w http.ResponseWriter, r *http.Request) error {
	log.Println("nom nom")

	var ev Event
	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		return err
	}
	log.Println(ev)

	// Return it back during testing,.
	return json.NewEncoder(w).Encode(ev)
}
