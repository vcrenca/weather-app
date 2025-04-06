package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"weather-api/internal/app"
	"weather-api/internal/common"
	httpserver "weather-api/internal/ports/http"
	apiv1 "weather-api/internal/ports/http/v1"
)

func init() {
	env := os.Getenv("ENV")
	common.InitLogger(env)
	loadDotEnvIfLocal(env)
	slog.Info(
		"Starting application",
		slog.String("env", env),
	)
}

func loadDotEnvIfLocal(env string) {
	if env == "local" {
		if err := godotenv.Load(); err != nil {
			slog.Error("failed to load configuration from .env file", slog.String("error", err.Error()))
		}
	}
}

func main() {
	application := app.New()
	config := application.Configuration()

	server := httpserver.CreateServer(config.HttpPort, apiv1.CreateWeatherHttpHandler)
	slog.Info("Starting HTTP server", slog.String("port", config.HttpPort))
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start http server", slog.String("error", err.Error()))
		}
	}
}
