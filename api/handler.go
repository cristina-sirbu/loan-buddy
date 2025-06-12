package api

import (
	"encoding/json"
	"net/http"

	"github.com/cristina-sirbu/loan-buddy/internal/aggregator"
)

func OffersHandler(w http.ResponseWriter, r *http.Request) {
	offers := aggregator.AggregateOffers()
	w.Header().Set("Content-Type", "application/json")
	if len(offers) == 0 {
		http.Error(w, "No offers available", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(offers); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
