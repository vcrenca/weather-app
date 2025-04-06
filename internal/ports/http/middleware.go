package http

import (
	"log/slog"
	"net/http"
	"time"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		slog.Info(
			"request completed",
			slog.Group(
				"request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.String("duration", time.Since(start).String()),
			))
	})
}
