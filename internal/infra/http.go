package infra

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

const (
	requestIDKey = "requestID"
)

func CreateRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/toto", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World !"))
	}))

	return requestID(logging(mux))
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestID := mustGetRequestID(req)
		logger := slog.With(
			slog.String(requestIDKey, requestID),
			slog.String("method", req.Method),
			slog.String("path", req.URL.Path),
		)

		wrappedResponseWriter := newStatusCodeResponseWriter(w)
		next.ServeHTTP(wrappedResponseWriter, req)

		status := wrappedResponseWriter.status
		log := logger.Info
		if status >= 400 {
			log = logger.Error
		}

		log("request completed", "status", wrappedResponseWriter.status)
	})
}

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestID := uuid.New().String()

		ctx := context.WithValue(req.Context(), requestIDKey, requestID)
		req = req.WithContext(ctx)

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, req)
	})
}

func mustGetRequestID(req *http.Request) string {
	if requestID, ok := req.Context().Value(requestIDKey).(string); ok {
		return requestID
	}

	panic(fmt.Sprintf("no %s is set", requestIDKey))
}

type statusCodeResponseWriter struct {
	responseWriter http.ResponseWriter
	status         int
}

func newStatusCodeResponseWriter(w http.ResponseWriter) *statusCodeResponseWriter {
	return &statusCodeResponseWriter{
		responseWriter: w,
		status:         0,
	}
}

func (w *statusCodeResponseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *statusCodeResponseWriter) Write(bytes []byte) (int, error) {
	return w.responseWriter.Write(bytes)
}

func (w *statusCodeResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.responseWriter.WriteHeader(statusCode)
}