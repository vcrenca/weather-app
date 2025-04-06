package app

type Application struct {
	configuration Configuration
}

func (a Application) Configuration() Configuration {
	return a.configuration
}

func New() Application {
	return Application{
		configuration: getConfiguration(),
	}
}
