package weather

type Percent int

func NewPercent(value int) Percent {
	if value <= 0 || value >= 100 {
		panic("percent value can not be out of [0, 100].")
	}

	return Percent(value)
}

func (p Percent) Int() int {
	return int(p)
}

type Current struct {
	City               string
	Description        string
	TemperatureCelsius int
	WindKmPerHour      int
	RelativeHumidity   Percent
}
