package function

import (
	"fmt"
	"net/http"

	conf "example.com/module/internal/common/conf"
	routesHello "example.com/module/internal/example_hello_domain/http"
	routesVotes "example.com/module/internal/votes/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gin-gonic/gin"
)

var routerV1 *gin.Engine

func init() {
	routerV1 = gin.Default()

	err := conf.ConnectToMongoDB()
	if err != nil {
		panic(err)
	}

	// Configurar dependencias
	deps := routesVotes.NewAppDependencies()

	// Routes Group by Domain
	routesHello.Routes(routerV1)
	routesVotes.Routes(routerV1, deps)

	functions.HTTP("hello", Hello)
	routerV1.Run(":8080")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling func hello")
	routerV1.ServeHTTP(w, r)
}
