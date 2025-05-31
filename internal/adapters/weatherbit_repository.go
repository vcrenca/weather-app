package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"weather-api/internal/domain/weather"
)

type WeatherBitRepository struct {
	baseUrl string
	apiKey  string
}

func NewWeatherBitRepository(baseUrl string, apiKey string) *WeatherBitRepository {
	if baseUrl == "" {
		panic("missing baseUrl")
	}

	if apiKey == "" {
		panic("missing apiKey")
	}

	return &WeatherBitRepository{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

func (w WeatherBitRepository) GetCurrentWeather(ctx context.Context, city string) (*weather.Current, error) {
	cityQuery := make(url.Values, 1)
	cityQuery.Set("city", city)

	_, body, err := w.get(ctx, "/current", cityQuery)
	if err != nil {
		return nil, err
	}

	var weatherCurrentResponse WeatherBitCurrentResponse
	if err := json.Unmarshal(body, &weatherCurrentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode current weather response: %w", err)
	}

	if len(weatherCurrentResponse.Data) == 0 {
		return nil, weather.ErrWeatherNotFound
	}

	return mapToDomain(weatherCurrentResponse.Data[0]), nil
}

func (w WeatherBitRepository) get(ctx context.Context, path string, params url.Values) (*int, []byte, error) {
	requestURL, err := w.buildUrl(path, params)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build requestURL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &resp.StatusCode, responseBody, fmt.Errorf("WeatherBit responded with unexpected status code %d with body %s", resp.StatusCode, responseBody)
	}

	return &resp.StatusCode, responseBody, nil
}

func (w WeatherBitRepository) buildUrl(endpoint string, queryParams url.Values) (*url.URL, error) {
	urlPath, err := url.JoinPath(w.baseUrl, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join paths in url: %w", err)
	}

	requestURL, err := url.ParseRequestURI(urlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse path into url: %w", err)
	}

	queryParams.Add("key", w.apiKey)
	requestURL.RawQuery = queryParams.Encode()

	return requestURL, nil
}

func (w WeatherBitRepository) addApiKeyToUrl(baseURL *url.URL) *url.URL {
	urlCopy := *baseURL
	query := urlCopy.Query()
	query.Set("key", w.apiKey)
	urlCopy.RawQuery = query.Encode()
	return &urlCopy
}

type WeatherBitCurrentResponse struct {
	Data []WeatherBitCurrentResponseData `json:"data"`
}

type WeatherBitCurrentResponseData struct {
	CityName string  `json:"city_name"`
	Temp     float64 `json:"temp"`
	WindSpd  float64 `json:"wind_spd"`
	Rh       int     `json:"rh"`
	Weather  struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func mapToDomain(weatherBitResponse WeatherBitCurrentResponseData) *weather.Current {
	return &weather.Current{
		City:               weatherBitResponse.CityName,
		Description:        weatherBitResponse.Weather.Description,
		TemperatureCelsius: int(math.Round(weatherBitResponse.Temp)),
		WindKmPerHour:      int(math.Round(weatherBitResponse.WindSpd * 3.6)),
		RelativeHumidity:   weather.NewPercent(weatherBitResponse.Rh),
	}
}
