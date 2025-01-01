// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/gin-api/controllers"
)

// SetupRouter sets up the application routes
func SetupRouter(r *gin.Engine) {
	r.GET("/api/users", controllers.GetUsers)
	r.POST("/api/users", controllers.CreateUser)
}
