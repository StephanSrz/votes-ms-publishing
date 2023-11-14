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

	err := conf.ConnectToMongoDB(app.Env.DBHost, app.Env.DBUser, app.Env.DBPass, app.Env.DBCluster)
	if err != nil {
		panic(err)
	}
	app.Deps = votesHttp.NewAppDependencies()

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
