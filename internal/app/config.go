package app

import "os"

type Configuration struct {
	HttpPort string
}

func getConfiguration() Configuration {
	return Configuration{
		HttpPort: getOrDefault(os.Getenv("HTTP_PORT"), "8080"),
	}
}

func getOrDefault(value string, fallback string) string {
	if value == "" {
		return fallback
	}

	return value
}