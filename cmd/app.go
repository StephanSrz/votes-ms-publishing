package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

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
	// Configurar entorno
	var app = &App{}
	// Obtener la ruta del directorio actual
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get current file path")
	}
	baseDir := filepath.Dir(currentFile)

	// Construir la ruta completa al archivo .env.yaml
	configFilePath := filepath.Join(baseDir, ".env.yaml")
	fmt.Println(configFilePath)
	app.Env = conf.NewEnv(configFilePath)

	err := conf.ConnectToMongoDB(app.Env.DBHost, app.Env.DBUser, app.Env.DBPass, app.Env.DBCluster)
	if err != nil {
		panic(err)
	}
	// Configurar dependencias de la aplicación
	app.Deps = votesHttp.NewAppDependencies()

	// Configurar el enrutador
	app.Router = gin.Default()

	// Configurar rutas
	exampleHttp.Routes(app.Router)
	votesHttp.Routes(app.Router, app.Deps)

	return app
}

func (app *App) Start() {
	// Iniciar el servidor en la dirección especificada en la configuración
	addr := fmt.Sprintf("%s:%s", app.Env.ServerAddress, app.Env.DBPort)
	log.Printf("Server is running on %s", addr)
	app.Router.Run(addr)
}
