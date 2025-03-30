package http

import (
	"weather-api/internal/app"
	pkghttp "weather-api/pkg/http"
)

func RegisterHttpRoutes(application app.Application) *pkghttp.Router {
	router := pkghttp.NewRouter()
	router.AddGroup("/api/v1/weather", WeatherRoutes(application))

	return router
}