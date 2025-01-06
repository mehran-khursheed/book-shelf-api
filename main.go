package main

import (
	"example/GIN-PROJECT/db"
	"example/GIN-PROJECT/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.ConnectDatabase("mongodb+srv://mehrankhursheed123:Mehranz1234@medipulse.j8y4z.mongodb.net/", "bookShelf")
	db.InitializeCollections("books")

	// Set up Gin router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Book API!",
		})
	})

	// Register routes
	routers.SetupRoutes(router)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}
