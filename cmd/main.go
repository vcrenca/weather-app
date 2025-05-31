package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"weather-api/internal/app"
	"weather-api/internal/common"
	server "weather-api/internal/ports/http"
	apiv1 "weather-api/internal/ports/http/v1"
)

func init() {
	env := os.Getenv("ENV")
	common.InitLogger(env)

	if env == app.LocalEnv {
		if err := godotenv.Load(); err != nil {
			slog.Error("failed to load configuration from .env file", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}

	slog.Info(
		"Starting application",
		slog.String("env", env),
	)
}

func main() {
	a := app.New()

	srv := server.CreateServer(
		a.Config().HttpPort,
		apiv1.CreateWeatherHttpHandler(a.WeatherRepository),
	)

	slog.Info("Starting HTTP server", slog.String("port", a.Config().HttpPort))
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start http server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}
}
