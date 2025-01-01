// services/database.go

package services

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
// Why a pointer?
// In Go, using a pointer (*mongo.Database) rather 
// than a value (mongo.Database) means that 
// operations can be performed on the same object 
// throughout the program, rather than copying the
//  object each time it’s used. This is important 
//  for large objects like a MongoDB connection
//   because it allows the program to save on 
//   memory and improve performance
var DB *mongo.Database

func ConnectDB(uri, dbName string) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database(dbName)
    log.Println("Connected to MongoDB!")
}
