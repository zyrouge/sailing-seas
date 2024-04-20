package core

type App struct {
	Env *Env
}

func CreateApp() (*App, error) {
	env, err := ParseEnv()
	if err != nil {
		return nil, err
	}
	app := &App{
		Env: env,
	}
	return app, nil
}
