package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"weather-api/internal/app"
	"weather-api/internal/infra"
)

func main() {
	config := app.GetConfig()

	router := infra.CreateRouter()

	slog.Info("Starting http server on port " + config.HttpPort)
	if err := http.ListenAndServe(":"+config.HttpPort, router); err != nil {
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