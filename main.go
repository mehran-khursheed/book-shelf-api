package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mehran-khursheed/gin-project/config"
	"github.com/mehran-khursheed/gin-project/routes"
	"github.com/mehran-khursheed/gin-project/services"
)

func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	// Connect to MongoDB
	services.ConnectDB(cfg.MongoURI, cfg.DBName)

	// Create a Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRouter(r)

	// Run the server
	r.Run(":" + cfg.Port) // Listen and serve on the specified port
}
