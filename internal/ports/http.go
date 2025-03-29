package ports

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

type ApiResponse struct {
	Message string `json:"message"`
}

func ApiV1Mux() http.Handler {
	apiMux := http.NewServeMux()

	apiMux.HandleFunc("GET /weather", func(w http.ResponseWriter, req *http.Request) {
		sendJsonResponse(w, http.StatusOK, ApiResponse{Message: "Weather !"})
	})

	apiMux.HandleFunc("GET /forecast", func(w http.ResponseWriter, req *http.Request) {
		sendJsonResponse(w, http.StatusOK, ApiResponse{Message: "Forecast !"})
	})

	return http.StripPrefix("/api/v1", apiMux)
}