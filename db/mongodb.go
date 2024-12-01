package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

// ConnectMongoDB initializes a connection to the MongoDB server with TLS
func ConnectMongoDB() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}

// GetClient returns the MongoDB client instance (singleton)
func GetClient() *mongo.Client {
	return clientInstance
}
