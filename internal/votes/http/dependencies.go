package http

import (
	repository "example.com/module/internal/votes/repository/mongo"
	services "example.com/module/internal/votes/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppDependencies struct {
	VoteRepo    repository.VoteRepository
	VoteService services.VoteService
	VoteHandler VoteHandler
}

func NewAppDependencies(dbInstanceConn *mongo.Database) *AppDependencies {
	dbInstance := dbInstanceConn
	voteRepo := repository.NewVoteRepository(dbInstance)
	voteService := services.NewVoteService(voteRepo)
	voteHandler := NewVoteHandler(voteService)

	return &AppDependencies{
		VoteRepo:    voteRepo,
		VoteService: voteService,
		VoteHandler: voteHandler,
	}
}
