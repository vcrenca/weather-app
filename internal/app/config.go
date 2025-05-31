package app

import (
	"fmt"
	"log/slog"
	"os"
)

const LocalEnv = "LOCAL"

type Configuration struct {
	HttpPort string

	WeatherBitBaseURL string
	WeatherBitAPIKey  string
}

func loadConfiguration() Configuration {
	return Configuration{
		HttpPort:          getOrDefault(os.Getenv("HTTP_PORT"), "8080"),
		WeatherBitBaseURL: getOrDefault(os.Getenv("WEATHER_BIT_BASE_URL"), "https://api.weatherbit.io/v2.0"),
		WeatherBitAPIKey:  require("WEATHER_BIT_API_KEY"),
	}
}

func getOrDefault(value string, fallback string) string {
	if value == "" {
		return fallback
	}

	return value
}

func require(variableName string) string {
	value := os.Getenv(variableName)
	if value == "" {
		panic(fmt.Sprintf("Environment variable %s not set", variableName))
	}

	return value
}

func (c Configuration) LogValue() slog.Value {
	attributes := []slog.Attr{
		slog.String("http_port", c.HttpPort),
		slog.String("weather_bit_base_url", c.WeatherBitBaseURL),
	}

	return slog.GroupValue(attributes...)
}
