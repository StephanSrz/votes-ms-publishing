package cmd

import (
	"fmt"
	"log"

	conf "example.com/module/internal/common/conf"
	votesHttp "example.com/module/internal/votes/http"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router            *gin.Engine
	VotesDependencies *votesHttp.AppDependencies
	Env               *conf.Env
}

func NewApp() *App {
	var app = &App{}

	app.Env = conf.NewEnv()

	app.VotesDependencies = votesHttp.NewAppDependencies(app.Env)

	app.Router = gin.Default()

	votesHttp.Routes(app.Router, app.VotesDependencies)

	return app
}

func (app *App) Start() {
	addr := fmt.Sprintf("http://localhost:%s", app.Env.ServerAddress)
	log.Printf("Server is running on %s", addr)
	app.Router.Run(app.Env.PortServer)
}
