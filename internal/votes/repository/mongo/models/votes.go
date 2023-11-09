package models

type Vote struct {
	UserId    string `json:"UserId" bson:"UserId"`
	RoomId    string `json:"RoomId" bson:"RoomId"`
	UserStory string `json:"UserStory" bson:"UserStory"`
	Value     string `json:"value" bson:"value"`
}