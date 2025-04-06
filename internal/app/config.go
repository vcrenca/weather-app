package app

import (
	"log/slog"
	"os"
)

type Configuration struct {
	HttpPort string

	WeatherBitBaseURL string
	WeatherBitAPIKey  string
}

func getConfiguration() Configuration {
	return Configuration{
		HttpPort:          getOrDefault(os.Getenv("HTTP_PORT"), "8080"),
		WeatherBitBaseURL: getOrDefault(os.Getenv("WEATHER_BIT_BASE_URL"), ""),
		WeatherBitAPIKey:  getOrDefault(os.Getenv("WEATHER_BIT_API_KEY"), ""),
	}
}

func getOrDefault(value string, fallback string) string {
	if value == "" {
		return fallback
	}

	return value
}

func (c Configuration) LogValue() slog.Value {
	attributes := []slog.Attr{
		slog.String("http_port", c.HttpPort),
		slog.String("weather_bit_base_url", c.WeatherBitBaseURL),
	}

	if os.Getenv("ENV") == "local" {
		attributes = append(attributes, slog.String("weather_bit_api_key", c.WeatherBitAPIKey))

	}

	return slog.GroupValue(attributes...)
}
