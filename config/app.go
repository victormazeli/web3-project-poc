package config

type Application struct {
	Env *Env
}

func LoadEnvironmentVariables() Application {
	app := &Application{}
	app.Env = NewEnv()
	return *app
}
