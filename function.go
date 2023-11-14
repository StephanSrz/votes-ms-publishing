package function

import (
	"fmt"
	"net/http"

	cmd "example.com/module/cmd"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var app *cmd.App

func init() {

	app := cmd.NewApp()

	app.Start()

	// Configurar funciones HTTP
	functions.HTTP("hello", Hello )
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling func hello")
	app.Router.ServeHTTP(w, r)
}
