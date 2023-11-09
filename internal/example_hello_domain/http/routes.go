package http

import (
	"github.com/gin-gonic/gin"
)

// c *gin.Context
func Routes(route *gin.Engine) {
	routeHello := route.Group("/hello")
	routeHello.GET("/ping", Hello)
	routeHello.POST("/ping", HelloName)
}
