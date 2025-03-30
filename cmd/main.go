package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
	"weather-api/internal/app"
	portshttp "weather-api/internal/ports/http"
	pkghttp "weather-api/pkg/http"

	// Init logging configuration
	_ "weather-api/pkg/log"
)

func main() {
	application := app.New()
	config := application.Configuration()

	httpRouter := portshttp.RegisterHttpRoutes(application)
	server := pkghttp.CreateServer(config.HttpPort, httpRouter)
	slog.Info("Starting http server on port " + config.HttpPort)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start http server", "err", err)
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