package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Event struct {
	Type    int `json:"type"`
	Payload any `json:"payload"`
}

type Response struct {
	Data  any   `json:"data"`
	Error error `json:"error"`
}

// Ingest sounds like a cool name for whatever this is going to end up being. Definitely not for a rule engine tho :(
func (h *Handler) Ingest(w http.ResponseWriter, r *http.Request) error {
	log.Println("nom nom")

	var evs []Event
	if err := json.NewDecoder(r.Body).Decode(&evs); err != nil {
		return err
	}

	resps := make(chan Response, len(evs))

	// Move workers into somewhere else like dispatch maybe?
	var wg sync.WaitGroup
	wg.Add(len(evs))

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	for _, ev := range evs {
		go func(ctx context.Context, resps chan Response) {
			defer wg.Done()

			var resp Response

			service, err := h.ss.Fetch(ev.Type)
			if err != nil {
				resp.Error = err
				resps <- resp
				return
			}
			// Dispatch will handle converting into the correct payload format for the service.
			data, err := h.d.Dispatch(ctx, service, ev.Payload)
			if err != nil {
				resp.Error = err
				resps <- resp
				return
			}

			resp.Data = data

			resps <- resp
		}(ctx, resps)
	}

	go func() {
		wg.Wait()
		close(resps)
	}()

	respReturns := []Response{}

	for resp := range resps {
		respReturns = append(respReturns, resp)
	}

	return json.NewEncoder(w).Encode(respReturns)
}
