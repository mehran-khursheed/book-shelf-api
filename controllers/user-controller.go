package controllers

import (
    "context"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/yourusername/gin-api/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/yourusername/gin-api/services"
)

var userCollection = services.DB.Collection("users")

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.ID = primitive.NewObjectID()
    user.CreatedAt = time.Now().Format(time.RFC3339)
    user.UpdatedAt = time.Now().Format(time.RFC3339)

    _, err := userCollection.InsertOne(context.Background(), user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
    var users []models.User
    cursor, err := userCollection.Find(context.Background(), bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
        return
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }

    c.JSON(http.StatusOK, users)
}
