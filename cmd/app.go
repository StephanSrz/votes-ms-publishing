package cmd

import (
	conf "example.com/module/internal/common/conf"
)

type Application struct {
	Env *conf.Env
}

func App() *Application {
	app := &Application{}
	app.Env = conf.NewEnv()

	return app
}

// func (app *Application) CloseDBConnection() {
// 	conf.CloseMongoDBConnection(app.Mongo)
// }
