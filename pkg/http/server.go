package http

import (
	"net/http"
)

func CreateServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: WithHttpRequestID(WithHttpLogging(handler)),
	}
}