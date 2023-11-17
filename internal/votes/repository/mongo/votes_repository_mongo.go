package repository

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	entity "example.com/module/internal/votes/entity"
	models "example.com/module/internal/votes/repository/mongo/models"
	"github.com/spf13/viper"
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
	collection := vr.db.Collection("votes_test")

	ctx := context.TODO()
	fmt.Println("Insertando en la base de datos...")
	_, err := collection.InsertOne(ctx, voteModel)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserción exitosa.")

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
