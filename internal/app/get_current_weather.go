package app

import "weather-api/internal/domain/weather"

type GetCurrentWeatherByCity func(city string) (weather.Current, error)