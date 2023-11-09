package conf

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectToMongoDB() error {
	mongoURI := "" // MONGO DB ATLAS URL
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	MongoClient = client
	return nil
}
