package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, deps *AppDependencies) {
	routesVotes := router.Group("/v1/votes")
	routesVotes.POST("", deps.VoteHandler.CreateVote)
}
