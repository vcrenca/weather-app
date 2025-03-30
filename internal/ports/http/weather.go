package http

import (
	"net/http"
	"weather-api/internal/app"
	pkghttp "weather-api/pkg/http"
)

func WeatherRoutes(application app.Application) http.Handler {
	router := pkghttp.NewRouter()
	router.Handle("/current", GetCurrentWeather(application))
	router.Handle("/forecast", GetForecast(application))
	return router
}

func GetCurrentWeather(app app.Application) pkghttp.Handler {
	return func(w http.ResponseWriter, req *http.Request) error {
		w.WriteHeader(http.StatusOK)

		return nil
	}
}

func GetForecast(app app.Application) pkghttp.Handler {
	return func(w http.ResponseWriter, req *http.Request) error {
		w.WriteHeader(http.StatusAccepted)

		return nil
	}
}