package db

import (
	"context"
	"eventapp/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context

// ConnectDB initializes the MongoDB client and sets the context.
func ConnectDB() (*mongo.Client, error) {
	appConfig := config.LoadConfig()

	var err error
	ctx = context.TODO() // Set the context for the application
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(appConfig.MONGO_URI))
	if err != nil {
		log.Println("Error on connection time MONGO_URI:", err)
		return client, err
	}
	log.Println("Database Connection Done...")

	// Optionally ping the database to ensure the connection is established
	if err = client.Ping(ctx, nil); err != nil {
		log.Println("Error on connection Pinging Time:", err)
		return client, err
	}
	log.Println("Database Connection Pinging Done...")

	return client, nil
}

// // GetDatabase returns the database instance.
// func GetDatabase() (*mongo.Database, context.Context) {
// 	appConfig := config.LoadConfig()
// 	return client.Database(appConfig.DATABASENAME), ctx
// }

// // DisconnectDB closes the MongoDB connection.
// func DisconnectDB() error {
// 	if client != nil {
// 		return client.Disconnect(ctx)
// 	}
// 	return nil
// }
