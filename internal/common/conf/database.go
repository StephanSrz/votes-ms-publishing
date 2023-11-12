package conf

// import (
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var (
// 	//	client *mongo.Client
// 	db *mongo.Database
// )

// func ConnectDB(env *Env) error {

// 	dbHost := env.DBHost
// 	dbUser := env.DBUser
// 	dbPass := env.DBPass
// 	dbCluster := env.DBCluster

// 	mongodbURI := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority", dbHost, dbUser, dbPass, dbCluster)
// 	fmt.Print(mongodbURI)
// 	clientOptions := options.Client().ApplyURI(mongodbURI)
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		return err
// 	}

// 	// Comprueba la conexión
// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		return err
// 	}

// 	db = client.Database("database_name")
// 	fmt.Println("Connected to MongoDB")
// 	return nil
// }

// func GetDB() *mongo.Database {
// 	return db
// }
