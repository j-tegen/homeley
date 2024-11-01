package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
    Client *mongo.Client
}

// NewDatabase initializes the Database struct with a MongoDB client
func NewDatabase(uri string) (*Database, error) {
    client, err := connectMongo(uri)
    if err != nil {
        return nil, err
    }
    return &Database{Client: client}, nil
}

func connectMongo(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
			return nil, err
	}

	// Ping the database to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	if err := client.Ping(ctx, nil); err != nil {
			return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}
