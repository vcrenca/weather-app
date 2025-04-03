package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func CreateWeatherHttpHandler(router chi.Router) http.Handler {
	return HandlerFromMux(WeatherHttpHandler{}, router)
}

type WeatherHttpHandler struct{}

func (s WeatherHttpHandler) GetV1WeatherCurrent(w http.ResponseWriter, r *http.Request, params GetV1WeatherCurrentParams) {
	render.Respond(w, r, CurrentWeather{
		Description: "The current weather for " + params.Location,
		Humidity:    56,
		Temperature: 31,
		Wind:        45.3,
	})
}

func (s WeatherHttpHandler) GetV1WeatherForecast(w http.ResponseWriter, r *http.Request, params GetV1WeatherForecastParams) {
	render.Respond(w, r, Forecast{
		AverageWind:      LightAir,
		GeneralTrend:     ForecastGeneralTrendDeteriorating,
		PressureTrend:    PressureTrendFalling,
		TemperatureTrend: Stable,
	})
}
