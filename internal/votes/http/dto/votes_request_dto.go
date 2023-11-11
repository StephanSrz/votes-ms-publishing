package dto

import (
	entity "example.com/module/internal/votes/entity"
)

type VotesRequestDTO struct {
	UserId    string `json:"UserId" validate:"required"`
	RoomId    string `json:"RoomId" validate:"required"`
	UserStory string `json:"UserStory" validated:"omitempty"`
	Value     string `json:"Value" validate:"required"`
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
