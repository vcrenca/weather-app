package app

import "os"

type Config struct {
	HttpPort string
}

func GetConfig() Config {
	return Config{
		HttpPort: getOrDefault(os.Getenv("HTTP_PORT"), "8080"),
	}
}

func getOrDefault(value string, fallback string) string {
	if value == "" {
		return fallback
	}

	return value
}