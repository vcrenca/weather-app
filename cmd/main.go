package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"weather-api/internal/adapters"
	"weather-api/internal/app"
	"weather-api/internal/common"
	server "weather-api/internal/ports/http"
	apiv1 "weather-api/internal/ports/http/v1"
)

func init() {
	env := os.Getenv("ENV")
	common.InitLogger(env)

	if env == "local" {
		if err := godotenv.Load(); err != nil {
			slog.Error("failed to load configuration from .env file", slog.String("error", err.Error()))
		}
	}

	slog.Info(
		"Starting application",
		slog.String("env", env),
	)
}

func main() {
	application := app.New()
	config := application.Configuration()

	slog.Info("Loaded configuration from .env file", "configuration", config)

	weatherBitRepository := adapters.NewWeatherBitRepository(
		config.WeatherBitBaseURL,
		config.WeatherBitAPIKey,
	)

	srv := server.CreateServer(config.HttpPort, apiv1.CreateWeatherHttpHandler(weatherBitRepository))
	slog.Info("Starting HTTP server", slog.String("port", config.HttpPort))
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start http server", slog.String("error", err.Error()))
		}
	}
}
