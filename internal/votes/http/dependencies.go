package http

import (
	conf "example.com/module/internal/common/conf"
	repository "example.com/module/internal/votes/repository/mongo"
	services "example.com/module/internal/votes/services"
)

type AppDependencies struct {
	VoteRepo    repository.VoteRepository
	VoteService services.VoteService
	VoteHandler VoteHandler
}

func NewAppDependencies(envVar *conf.Env) *AppDependencies {
	dbClient, _ := conf.ConnectToMongoDB(envVar.DBHost, envVar.DBUser, envVar.DBName, envVar.DBPass, envVar.DBCluster)
	voteRepo := repository.NewVoteRepository(dbClient)
	voteService := services.NewVoteService(voteRepo)
	voteHandler := NewVoteHandler(voteService)

	return &AppDependencies{
		VoteRepo:    voteRepo,
		VoteService: voteService,
		VoteHandler: voteHandler,
	}
}
