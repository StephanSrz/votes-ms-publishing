package repository

import (
	"context"
	"fmt"

	conf "example.com/module/internal/common/conf"
	entity "example.com/module/internal/votes/entity"
	models "example.com/module/internal/votes/repository/mongo/models"
)

type voteRepository struct{}

type VoteRepository interface {
    SaveVote(vote *entity.Votes) (*entity.Votes, error)
}

func NewVoteRepository() VoteRepository {
    return &voteRepository{}
}

func (vr *voteRepository) SaveVote(vote *entity.Votes) (*entity.Votes, error) {

	var voteModel models.Vote
	voteModel.MapFromEntity(vote)

	

	// Conectarse a MongoDB
	client := conf.MongoClient
	collection := client.Database("plannig-poker").Collection("votes")

	ctx := context.TODO()
	_, err := collection.InsertOne(ctx, voteModel)
	if err != nil {
		return nil, err
	}

	voteEntity := voteModel.MapToEntity()
	fmt.Println("---------------REPOSITORY----------------")
	fmt.Println(voteEntity)
	return voteEntity, nil
}
