package conf

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectToMongoDB(databaseData ...string) error {

	dbHost := databaseData[0]
	dbName := databaseData[1]
	DBPass := databaseData[2]
	dbCluster := databaseData[3]
	mongoURI := fmt.Sprintf("%s://%s:%s@%s", dbHost, dbName, DBPass, dbCluster) // MONGO DB ATLAS URL
	fmt.Println(mongoURI)
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
