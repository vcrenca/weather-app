package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"weather-api/internal/domain/weather"
)

type WeatherBitRepository struct {
	baseUrl string
	apiKey  string
}

// Response type for WeatherBit API
type WeatherBitCurrentResponse struct {
	CityName string  `json:"city_name"`
	Temp     float64 `json:"temp"`
	WindSpd  float64 `json:"wind_spd"`
	Rh       int     `json:"rh"`
	Weather  struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func mapToDomain(weatherBitResponse WeatherBitCurrentResponse) *weather.Current {
	return &weather.Current{
		City:               weatherBitResponse.CityName,
		Description:        weatherBitResponse.Weather.Description,
		TemperatureCelsius: int(weatherBitResponse.Temp),
		WindKmPerHour:      int(weatherBitResponse.WindSpd * 3.6),
		RelativeHumidity:   weather.NewPercent(weatherBitResponse.Rh),
	}
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
	response, err := w.get(ctx, city)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result struct {
		Data []WeatherBitCurrentResponse `json:"data"`
	}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result.Data) == 0 {
		//TODO: Implement business error ErrNoWeatherFound
		return nil, fmt.Errorf("no weather data returned")
	}

	return mapToDomain(result.Data[0]), nil
}

func (w WeatherBitRepository) get(ctx context.Context, endpoint string) (*http.Response, error) {
	url, err := w.buildUrl(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return resp, nil
}

func (w WeatherBitRepository) buildUrl(endpoint string) (*url.URL, error) {
	urlPath, err := url.JoinPath(w.baseUrl, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join paths in url: %w", err)
	}

	builtUrl, err := url.ParseRequestURI(urlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse path into url: %w", err)
	}

	return w.addApiKeyToUrl(builtUrl), nil
}

func (w WeatherBitRepository) addApiKeyToUrl(baseURL *url.URL) *url.URL {
	urlCopy := *baseURL
	query := urlCopy.Query()
	query.Set("key", w.apiKey)
	urlCopy.RawQuery = query.Encode()
	return &urlCopy
}
