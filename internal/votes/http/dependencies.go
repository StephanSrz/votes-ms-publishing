package http

import (
	repository "example.com/module/internal/votes/repository/mongo"
	services "example.com/module/internal/votes/services"
)

type AppDependencies struct {
	VoteRepo    repository.VoteRepository
	VoteService services.VoteService
	VoteHandler VoteHandler
}

func NewAppDependencies() *AppDependencies {
	voteRepo := repository.NewVoteRepository()
	voteService := services.NewVoteService(voteRepo)
	voteHandler := NewVoteHandler(voteService)

	return &AppDependencies{
		VoteRepo:    voteRepo,
		VoteService: voteService,
		VoteHandler: voteHandler,
	}
}
