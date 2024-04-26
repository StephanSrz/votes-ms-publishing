package dto

import (
	entity "example.com/module/internal/votes/entity"
)

type VotesRequestDTO struct {
	UserId    string `json:"userId" validate:"required"`
	RoomId    string `json:"roomId" validate:"required"`
	UserStory string `json:"userStory" validated:"omitempty"`
	Value     string `json:"value" validate:"required"`
}

func (dto *VotesRequestDTO) MapToVoteEntity() entity.Votes {

	entityVotes := entity.Votes{
		UserId:    dto.UserId,
		RoomId:    dto.RoomId,
		UserStory: dto.UserStory,
		Value:     dto.Value,
	}

	return entityVotes
}
