package http

import (
	"log/slog"
	"net/http"
)

type Handler func(w http.ResponseWriter, req *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := h(w, req); err != nil {
		slog.Error("error handling http request", "error", err.Error())
		InternalServerError(w, err.Error())
	}
}