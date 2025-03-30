package log

import (
	"log/slog"
	"os"
)

func init() {
	var logger *slog.Logger
	if os.Getenv("ENV") == "local" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	slog.SetDefault(logger)
}