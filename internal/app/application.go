package app

type Application struct {
	configuration Configuration

	GetCurrentWeatherByCity
}

func (a Application) Configuration() Configuration {
	return a.configuration
}

func New() Application {
	return Application{
		configuration: getConfiguration(),
	}
}