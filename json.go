package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type errResp struct {
	Error string `json:"error"`
}

type blogResp struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Category  string          `json:"category"`
	Tags      json.RawMessage `json:"tags"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}

func decodeJSON(req *http.Request, params any) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)
	return err
}
