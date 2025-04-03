package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	if os.Getenv("ENV") == "local" {
		log.SetFormatter(&log.TextFormatter{PadLevelText: true})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.SetLevel(log.DebugLevel)
}
