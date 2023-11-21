package http

import (
	"net/http"

	dtos "example.com/module/internal/votes/http/dto"
	services "example.com/module/internal/votes/services"
	"github.com/gin-gonic/gin"
)

type voteHandler struct {
	vs services.VoteService
}

type VoteHandler interface {
	CreateVote(c *gin.Context)
}

func NewVoteHandler(voteService services.VoteService) VoteHandler {
	return &voteHandler{
		vs: voteService,
	}
}

func (vh *voteHandler) CreateVote(c *gin.Context) {
	var voteDto dtos.VotesRequestDTO
	if err := c.ShouldBindJSON(&voteDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	voteEntity := voteDto.MapToVoteEntity()
	err := vh.vs.SaveVote(&voteEntity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Vote created"})
}

func GetVotes(c *gin.Context) {

}
