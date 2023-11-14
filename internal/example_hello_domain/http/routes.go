package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	routeHello := route.Group("/hello")
	routeHello.GET("/ping", Hello)
	routeHello.POST("/ping", HelloName)
}
