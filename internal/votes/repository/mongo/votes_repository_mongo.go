package repository

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	conf "example.com/module/internal/common/conf"
	entity "example.com/module/internal/votes/entity"
	models "example.com/module/internal/votes/repository/mongo/models"
	"github.com/spf13/viper"
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

	// Connecting to mongo
	client := conf.MongoClient
	collection := client.Database("planning-poker").Collection("votes")

	ctx := context.TODO()
	_, err := collection.InsertOne(ctx, voteModel)
	if err != nil {
		return nil, err
	}

	err = publishData(voteModel.UserId, voteModel.RoomId)
	if err != nil {
		panic(err)
	}

	voteEntity := voteModel.MapToEntity()
	fmt.Println("---------------REPOSITORY----------------")
	fmt.Println(voteEntity)
	return voteEntity, nil
}

func publishData(userId string, roomId string) error {
	pubToUrl := viper.GetString("PUBLISH_DATA_URL")

	bodyData := []byte(fmt.Sprintf(`{"UserId": "%s", "RoomId": "%s"}`, userId, roomId))

	req, err := http.NewRequest("POST", pubToUrl, bytes.NewBuffer(bodyData))
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return err
	}
	defer resp.Body.Close()
	return nil

}
