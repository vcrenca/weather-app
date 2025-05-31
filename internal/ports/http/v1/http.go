package v1

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"weather-api/internal/common"
	"weather-api/internal/domain/weather"
)

func CreateWeatherHttpHandler(repository weather.Repository) func(router chi.Router) http.Handler {
	return func(router chi.Router) http.Handler {
		return HandlerWithOptions(
			WeatherHttpHandler{weatherRepository: repository},
			ChiServerOptions{
				BaseRouter:       router,
				ErrorHandlerFunc: errHandler,
			},
		)
	}
}

type WeatherHttpHandler struct {
	weatherRepository weather.Repository
}

func (h WeatherHttpHandler) GetV1WeatherCurrent(w http.ResponseWriter, r *http.Request, params GetV1WeatherCurrentParams) {
	currentWeather, err := h.weatherRepository.GetCurrentWeather(r.Context(), params.City)
	if err != nil {
		if errors.Is(err, weather.ErrWeatherNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		slog.Error(
			"failed to get current weather",
			slog.String("city", params.City),
			slog.String("error", err.Error()),
		)
		internalServerError(w, r)
		return
	}

	render.Respond(w, r, CurrentWeather{
		Description:        currentWeather.Description,
		HumidityPercent:    common.Ptr(currentWeather.RelativeHumidity.ToInt()),
		TemperatureCelsius: common.Ptr(currentWeather.TemperatureCelsius),
		WindSpeedKmh:       common.Ptr(currentWeather.WindKmPerHour),
	})
}

func (h WeatherHttpHandler) GetV1WeatherForecast(w http.ResponseWriter, r *http.Request, params GetV1WeatherForecastParams) {
	render.Respond(w, r, Forecast{
		AverageWind:      LightAir,
		GeneralTrend:     ForecastGeneralTrendDeteriorating,
		PressureTrend:    PressureTrendFalling,
		TemperatureTrend: Stable,
	})
}
