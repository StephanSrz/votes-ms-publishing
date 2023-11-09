package dto

type VotesRequestDTO struct {
	UserId    string `json:"UserId" validate:"required"`
	RoomId    string `json:"RoomId" validate:"required"`
	UserStory string `json:"UserStory" validated:"omitempty"`
	Value     string `json:"Value" validate:"required"`
}
