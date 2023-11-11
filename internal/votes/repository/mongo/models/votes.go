package models

import (
	entity "example.com/module/internal/votes/entity"
)

type Vote struct {
	UserId    string `bson:"UserId"`
	RoomId    string `bson:"RoomId"`
	UserStory string `bson:"UserStory"`
	Value     string `bson:"value"`
}

func (model *Vote) MapFromEntity(voteEntity *entity.Votes) *Vote {

	model.UserId = voteEntity.UserId
	model.RoomId = voteEntity.RoomId
	model.UserStory = voteEntity.UserStory
	model.Value = voteEntity.Value
	return model
}

func (model *Vote) MapToEntity() *entity.Votes {

	voteEntity := entity.Votes{
		UserId:    model.UserId,
		RoomId:    model.RoomId,
		UserStory: model.UserStory,
		Value:     model.Value,
	}

	return &voteEntity
}
