package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var BookCollection *mongo.Collection

// ConnectDatabase initializes the database connection
func ConnectDatabase(uri, dbName string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	db = client.Database(dbName)
	log.Println("Connected to MongoDB")
}

// InitializeCollections initializes the collections you'll be using in the app
func InitializeCollections(collectionName string) {
	BookCollection = db.Collection(collectionName)
	log.Printf("Initialized collection: %s\n", collectionName)
}
