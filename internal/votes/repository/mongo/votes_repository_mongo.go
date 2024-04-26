package repository

import (
	"context"

	entity "example.com/module/internal/votes/entity"
	models "example.com/module/internal/votes/repository/mongo/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type voteRepository struct {
	db *mongo.Database
}

type VoteRepository interface {
	SaveVote(vote *entity.Votes) (*entity.Votes, error)
}

func NewVoteRepository(dbClient *mongo.Database) VoteRepository {
	return &voteRepository{
		db: dbClient,
	}
}

func (vr *voteRepository) SaveVote(vote *entity.Votes) (*entity.Votes, error) {

	var voteModel models.Vote
	voteModel.MapFromEntity(vote)

	// Connecting to mongo
	collection := vr.db.Collection("votes")
	ctx := context.TODO()
	_, err := collection.InsertOne(ctx, voteModel)
	if err != nil {
		return nil, err
	}

	voteEntity := voteModel.MapToEntity()
	return voteEntity, nil
}
