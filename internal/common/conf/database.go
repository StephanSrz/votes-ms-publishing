// package conf

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	// mongo "example.com/module/mongo"
// 	"go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
// )

// func NewMongoDatabase(env *Env) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	dbHost := env.DBHost
// 	dbPort := env.DBPort
// 	dbUser := env.DBUser
// 	dbPass := env.DBPass
// 	dbCluster := env.DBCluster

// 	mongodbURI := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority", dbHost, dbUser, dbPass, dbCluster)

// 	if dbUser == "" || dbPass == "" {
// 		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
// 	}

// 	client, err := mongo.Connect(context.Background(), mongodbURI)
//     if err != nil {
//         return err
//     }

//     // Comprueba la conexión
//     err = client.Ping(context.Background(), nil)
//     if err != nil {
//         return err
//     }

//     db = client.Database("database_name")
//     fmt.Println("Connected to MongoDB")
//     return nil
// }

// func CloseMongoDBConnection(client mongo.Client) {
// 	if client == nil {
// 		return
// 	}

// 	err := client.Disconnect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Connection to MongoDB closed.")
// }

// database.go

package conf

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

// ConnectDB establece la conexión a la base de datos MongoDB
func ConnectDB(env *Env) error {

	dbHost := env.DBHost
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbCluster := env.DBCluster

	mongodbURI := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority", dbHost, dbUser, dbPass, dbCluster)
	fmt.Print(mongodbURI)
	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Comprueba la conexión
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	db = client.Database("database_name")
	fmt.Println("Connected to MongoDB")
	return nil
}

// GetDB retorna la instancia de la base de datos
func GetDB() *mongo.Database {
	return db
}
