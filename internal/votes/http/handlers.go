package http

import (
	"context"
	"net/http"

	conf "example.com/module/internal/common/conf"
	dtos "example.com/module/internal/votes/http/dto"
	"github.com/gin-gonic/gin"
)

func CreateVote(c *gin.Context) {
	var vote dtos.VotesRequestDTO
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Conectarse a MongoDB
	client := conf.MongoClient
	collection := client.Database("plannig-poker").Collection("votes")

	ctx := context.TODO()
	_, err := collection.InsertOne(ctx, vote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Vote created"})
}

func GetVotes(c *gin.Context) {

}
