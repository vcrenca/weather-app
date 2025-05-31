package app

import (
	"weather-api/internal/adapters"
	"weather-api/internal/domain/weather"
)

type Application struct {
	configuration Configuration

	WeatherRepository weather.Repository
}

func New() Application {
	configuration := loadConfiguration()

	weatherBitRepository := adapters.NewWeatherBitRepository(
		configuration.WeatherBitBaseURL,
		configuration.WeatherBitAPIKey,
	)

	return Application{
		configuration: configuration,

		WeatherRepository: weatherBitRepository,
	}
}

func (a Application) Config() Configuration {
	return a.configuration
}
