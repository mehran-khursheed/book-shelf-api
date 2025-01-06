package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"example/GIN-PROJECT/db"
	"example/GIN-PROJECT/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertBook adds a new book to the database.
func InsertBook(c *gin.Context) {
	var book models.Book

	// Bind incoming JSON to the book struct
	if err := c.ShouldBindJSON(&book); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Insert the book into the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.BookCollection.InsertOne(ctx, book)
	if err != nil {
		log.Println("Error inserting book:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert book"})
		return
	}

	// Respond with the inserted ID
	c.JSON(http.StatusOK, gin.H{
		"message": "Book inserted successfully",
		"id":      result.InsertedID,
	})
}

// UpdateBook modifies an existing book's details.
func UpdateBook(c *gin.Context) {
	// Get the book ID from the URL parameter
	bookId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		log.Println("Invalid book ID:", bookId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Bind the JSON payload to a map for flexibility
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update the book in the database
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateData}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.BookCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating book:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	// Respond with the update result
	c.JSON(http.StatusOK, gin.H{
		"message":       "Book updated successfully",
		"modifiedCount": result.ModifiedCount,
	})
}

// DeleteBook removes a book from the database.
func DeleteBook(c *gin.Context) {
	// Get the book ID from the URL parameter
	bookId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		log.Println("Invalid book ID:", bookId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Delete the book from the database
	filter := bson.M{"_id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.BookCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting book:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	// Respond with the deletion result
	c.JSON(http.StatusOK, gin.H{
		"message":      "Book deleted successfully",
		"deletedCount": result.DeletedCount,
	})
}

// GetAllBooks fetches all books from the database.
func GetAllBooks(c *gin.Context) {
	// Retrieve all books from the collection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := db.BookCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println("Error fetching books:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer cur.Close(ctx)

	// Decode books into a slice
	var books []bson.M
	for cur.Next(ctx) {
		var book bson.M
		if err := cur.Decode(&book); err != nil {
			log.Println("Error decoding book:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
			return
		}
		books = append(books, book)
	}

	// Respond with the retrieved books
	c.JSON(http.StatusOK, gin.H{"books": books})
}
