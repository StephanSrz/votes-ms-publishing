package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routesVotes := router.Group("/votes")
	routesVotes.GET("/", GetVotes)
	routesVotes.POST("/votes", CreateVote)
}
