package services

import (
	"fmt"

	entity "example.com/module/internal/votes/entity"
	repository "example.com/module/internal/votes/repository/mongo"
)

type voteService struct {
	repo repository.VoteRepository
}

type VoteService interface {
	SaveVote(entityVote *entity.Votes) error
}

func NewVoteService(repo repository.VoteRepository) VoteService {
	return &voteService{
		repo: repo,
	}
}

func (vs *voteService) SaveVote(entityVote *entity.Votes) error {

	fmt.Println("---------------SERVICE----------------")
	fmt.Println(entityVote)
	_, err := vs.repo.SaveVote(entityVote)
	if err != nil {
		return err
	}

	return nil
}
