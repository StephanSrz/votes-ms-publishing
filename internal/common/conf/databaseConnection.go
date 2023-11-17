package conf

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectToMongoDB(databaseData ...string) (*mongo.Database, error) {

	fmt.Println(databaseData)
	dbHost := databaseData[0]
	dbUser := databaseData[1]
	dbName := databaseData[2]
	DBPass := databaseData[3]
	dbCluster := databaseData[4]
	mongoURI := fmt.Sprintf("%s://%s:%s@%s", dbHost, dbUser, DBPass, dbCluster) // MONGO DB ATLAS URL

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	MongoClient = client
	return client.Database(dbName), nil
}
