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

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	MongoClient = client
	return nil
}
