package cmd

import (
	"fmt"
	"log"

	conf "example.com/module/internal/common/conf"
	exampleHttp "example.com/module/internal/example_hello_domain/http"
	votesHttp "example.com/module/internal/votes/http"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Deps   *votesHttp.AppDependencies
	Env    *conf.Env
}

func NewApp() *App {
	var app = &App{}

	app.Env = conf.NewEnv()

	app.Deps = votesHttp.NewAppDependencies(app.Env)

	app.Router = gin.Default()

	exampleHttp.Routes(app.Router)
	votesHttp.Routes(app.Router, app.Deps)

	return app
}

func (app *App) Start() {
	addr := fmt.Sprintf("http://localhost:%s", app.Env.ServerAddress)
	log.Printf("Server is running on %s", addr)
	app.Router.Run(app.Env.PortServer)
}
