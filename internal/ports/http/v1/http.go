package v1

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"weather-api/internal/domain/weather"
)

func CreateWeatherHttpHandler(repository weather.Repository) func(router chi.Router) http.Handler {
	return func(router chi.Router) http.Handler {
		return HandlerFromMux(
			WeatherHttpHandler{weatherRepository: repository},
			router,
		)
	}
}

type WeatherHttpHandler struct {
	weatherRepository weather.Repository
}

func (s WeatherHttpHandler) GetV1WeatherCurrent(w http.ResponseWriter, r *http.Request, params GetV1WeatherCurrentParams) {
	currentWeather, err := s.weatherRepository.GetCurrentWeather(r.Context(), params.City)
	if err != nil {
		if errors.Is(err, weather.ErrWeatherNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		slog.Error(
			"failed to get current weather",
			slog.String("city", params.City),
			slog.String("error", err.Error()),
		)
		return
	}

	render.Respond(w, r, CurrentWeather{
		Description: currentWeather.Description,
		Humidity:    currentWeather.RelativeHumidity.Int(),
		Temperature: currentWeather.TemperatureCelsius,
		Wind:        currentWeather.WindKmPerHour,
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
