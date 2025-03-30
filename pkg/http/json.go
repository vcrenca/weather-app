package http

import (
	"encoding/json"
	"net/http"
)

func sendJsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic("failed to write json data to http response: %s")
	}
}