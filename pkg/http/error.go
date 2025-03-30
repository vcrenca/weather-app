package http

import (
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func InternalServerError(w http.ResponseWriter, details string) {
	sendJsonResponse(w, http.StatusInternalServerError, ErrorResponse{
		Message: "Internal server error",
		Details: details,
	})
}