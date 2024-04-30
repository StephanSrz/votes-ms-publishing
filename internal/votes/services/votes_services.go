package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	entity "example.com/module/internal/votes/entity"
	repository "example.com/module/internal/votes/repository/mongo"
	"github.com/joho/godotenv"
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
	_, err := vs.repo.SaveVote(entityVote)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(`{"userID": "%s", "roomID": "%s", "userStory": "%s", "value": "%s"}`, entityVote.UserId, entityVote.RoomId, entityVote.UserStory, entityVote.Value)

	err = publishThatScales(os.Stdout, msg)
	if err != nil {
		panic(err)
	}

	return nil
}

var projectID string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectID = os.Getenv("PROJECT_ID")
}

func publishThatScales(w io.Writer, msg string) error {
	topicID := "votes"
	// msg := "Hello World"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %w", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
