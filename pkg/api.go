package deus_cc

import (
	"encoding/json"
	"log"
	"net/http"
)

type Event struct {
	Url  string
	Uuid string
}

type Handler struct {
	Storer *Storer
}

// SetEvent is the handler to manage ingestion of events
func (h Handler) SetEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			event := Event{}
			if err := json.NewDecoder(req.Body).Decode(&event); err != nil {
				responseUnprocessable(w, "")
				return
			}
			log.Printf("Inconming event: %v", event)

			if event.Url != "" && event.Uuid != "" {
				ValidateData(event)
				success := h.Storer.Adder(event)
				responseOk(w, success, 0)
				return
			}
			responseBadRequest(w)
			return
		}
		responseNotAllowed(w, req.Method)
	})
}

// GetDistinctVisitors is the handler to serve the number of distinct visitors of any given page
func (h Handler) GetDistinctVisitors() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			url := req.URL.Query().Get("url")
			if url == "" {
				responseError(w, "URL param not found")
				return
			}
			visitors := h.Storer.Getter(url)
			responseVisitors(w, true, visitors)
			return
		}
		responseNotAllowed(w, req.Method)
	})
}
