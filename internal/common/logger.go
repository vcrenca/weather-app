package common

import (
	"log/slog"
	"os"
)

func InitLogger(env string) {
	var logger *slog.Logger
	if env == "local" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(logger)
}
