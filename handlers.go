package main

import (
	"encoding/json"
	"net/http"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Generate unique ID
	id, err := idGen.NextID()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	shortID := encodeBase62(id)

	// Save mapping
	Save(shortID, req.URL)

	resp := shortenResponse {
		ShortURL: baseURL + shortID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[1:]
	longURL, ok := Get(shortID)

	if !ok {
		http.NotFound(w,r)
		return
	}

	http.Redirect(w,r, longURL, http.StatusFound)
}