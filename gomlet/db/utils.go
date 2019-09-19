package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectionString = ""
	dbName           = ""
	mongoClient      *mongo.Client
)

// Connect ...
func Connect(conn string, db string) {
	connectionString = conn
	dbName = db
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client

	// Check the connection
	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
