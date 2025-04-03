package http

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func LogrusMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Call the next handler in the chain
			next.ServeHTTP(w, r)

			// Log request details
			log.WithFields(log.Fields{
				"method":      r.Method,
				"url":         r.URL.String(),
				"remote_addr": r.RemoteAddr,
				"user_agent":  r.UserAgent(),
				"duration":    time.Since(start).String(),
			}).Info("request completed")
		})
	}
}
