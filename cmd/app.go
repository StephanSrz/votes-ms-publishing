package cmd

import (
	"fmt"
	"log"

	conf "example.com/module/internal/common/conf"
	votesHttp "example.com/module/internal/votes/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router            *gin.Engine
	VotesDependencies *votesHttp.AppDependencies
	DbConn            *mongo.Database
	Env               *conf.Env
}

func NewApp() *App {
	var app = &App{}

	app.Env = conf.NewEnv()

	Dbenv := &conf.DbEnv{
		DbEnviroment: app.Env.DbEnviroment,
		Server:       app.Env.MongoServer,
		Username:     app.Env.MongoUsername,
		Password:     app.Env.MongoPassword,
		Cluster:      app.Env.MongoCluster,
		Dbname:       app.Env.DbName,
	}

	app.DbConn = conf.GetDBInstance(Dbenv)

	app.VotesDependencies = votesHttp.NewAppDependencies(app.DbConn)

	app.Router = gin.Default()

	votesHttp.Routes(app.Router, app.VotesDependencies)

	return app
}

func (app *App) Start() {
	addr := fmt.Sprintf("http://localhost:%s", app.Env.ServerAddress)
	log.Printf("Server is running on %s", addr)
	app.Router.Run(app.Env.PortAddress)
}
