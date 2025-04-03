package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
	"weather-api/internal/app"
	httpserver "weather-api/internal/ports/http"
	apiv1 "weather-api/internal/ports/http/v1"

	// Init logging configuration
	_ "weather-api/pkg/log"
)

func main() {
	application := app.New()
	config := application.Configuration()

	server := httpserver.CreateServer(config.HttpPort, apiv1.CreateWeatherHandler())

	slog.Info("Starting HTTP server", "port", config.HttpPort)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start http server: %s", err.Error())
		}
	}
}

func init() {
	env := os.Getenv("ENV")
	slog.Info(
		"Starting application",
		"ENV", env,
	)

	if os.Getenv("ENV") == "local" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("failed to load configuration from .env file: %s", err)
		}
	}
}
