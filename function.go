package function

import (
	"fmt"
	"net/http"

	cmd "example.com/module/cmd"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {

	app := cmd.NewApp()

	app.Start()
	// Configurar funciones HTTP
	functions.HTTP("hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling func hello")
		app.Router.ServeHTTP(w, r)
	})
}

// func Hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Calling func hello")
// 	routerV1.ServeHTTP(w, r)
// }
