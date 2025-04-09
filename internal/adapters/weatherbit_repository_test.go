package adapters_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-api/internal/adapters"
	"weather-api/internal/domain/weather"
)

func mockGetCurrentWeatherByCity(city string, mock func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/current" && r.URL.Query().Get("city") == city {
			mock(w, r)
		}
	}

}

func TestWeatherBitRepository_GetCurrentWeather(t *testing.T) {
	t.Run("should return current weather", func(t *testing.T) {
		t.Parallel()

		// Given
		mockServer := httptest.NewServer(mockGetCurrentWeatherByCity("city", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{  
				   "data":[  
					  {  
						 "wind_cdir":"NE",
						 "rh":59,
						 "pod":"d",
						 "lon":-78.63861,
						 "pres":1006.6,
						 "timezone":"America\/New_York",
						 "ob_time":"2017-08-28 16:45",
						 "country_code":"US",
						 "clouds":75,
						 "vis":10,
						 "wind_spd":6.17,
						 "gust": 8,
						 "wind_cdir_full":"northeast",
						 "app_temp":24.25,
						 "state_code":"NC",
						 "ts":1503936000,
						 "h_angle":0,
						 "dewpt":15.65,
						 "weather":{  
							"icon":"c03d",
							"code": 803,
							"description":"Broken clouds"
						 },
						 "uv":2,
						 "aqi":45,
						 "station":"CMVN7",
						 "sources": ["rtma", "CMVN7"],
						 "wind_dir":50,
						 "elev_angle":63,
						 "datetime":"2017-08-28:17",
						 "precip":0,
						 "ghi":444.4,
						 "dni":500,
						 "dhi":120,
						 "solar_rad":350,
						 "city_name":"Raleigh",
						 "sunrise":"10:44",
						 "sunset":"23:47",
						 "temp":24.19,
						 "lat":35.7721,
						 "slp":1022.2
					  }
				   ],
				   "minutely":[],
				   "count":1
				}`))
		}))
		defer mockServer.Close()

		repository := adapters.NewWeatherBitRepository(mockServer.URL, "apiKey")

		// When
		result, err := repository.GetCurrentWeather(context.Background(), "city")

		// Then
		require.NoError(t, err)
		require.Equal(
			t,
			&weather.Current{
				City:               "Raleigh",
				Description:        "Broken clouds",
				TemperatureCelsius: 24,
				WindKmPerHour:      22,
				RelativeHumidity:   59,
			},
			result,
		)
	})

	t.Run("should return ErrWeatherNotFound when no data", func(t *testing.T) {
		t.Parallel()

		// Given
		mockResponse := `{"data":[]}`

		mockServer := httptest.NewServer(mockGetCurrentWeatherByCity("Paris", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(mockResponse))
		}))
		defer mockServer.Close()

		repo := adapters.NewWeatherBitRepository(mockServer.URL, "fake-api-key")

		// When
		_, err := repo.GetCurrentWeather(context.Background(), "Paris")

		// Then
		require.Error(t, err)
		assert.ErrorIs(t, err, weather.ErrWeatherNotFound)
	})

	t.Run("should return error when non-200 response", func(t *testing.T) {
		t.Parallel()

		// Given
		mockServer := httptest.NewServer(mockGetCurrentWeatherByCity("UnknownCity", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "city not found"}`))
		}))
		defer mockServer.Close()

		repo := adapters.NewWeatherBitRepository(mockServer.URL, "fake-api-key")

		// When
		_, err := repo.GetCurrentWeather(context.Background(), "UnknownCity")

		// Then
		require.Error(t, err)
		assert.Equal(t, "WeatherBit responded with unexpected status code 404 with body {\"error\": \"city not found\"}", err.Error())
	})
}
