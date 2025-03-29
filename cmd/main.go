package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"weather-api/internal/app"
	"weather-api/internal/infra"
	"weather-api/internal/ports"
)

func main() {
	config := app.GetConfig()

	rootMux := http.NewServeMux()
	rootMux.Handle("/api/v1/", ports.ApiV1Mux())

	server := infra.CreateServer(config.HttpPort, rootMux)
	slog.Info("Starting http server on port " + config.HttpPort)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start http server", "err", err)
		}
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load configuration from .env file: %s", err)
	}
}