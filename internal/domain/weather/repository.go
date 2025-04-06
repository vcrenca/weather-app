package weather

import "context"

type Repository interface {
	GetCurrentWeather(ctx context.Context, city string) (*Current, error)
}
