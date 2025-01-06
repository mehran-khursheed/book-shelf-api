package routers

import (
	"example/GIN-PROJECT/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	bookRoutes := router.Group("/books")
	{

		

		bookRoutes.POST("/post", controllers.InsertBook)
		bookRoutes.PUT("/update/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/delete/:id", controllers.DeleteBook)
		bookRoutes.GET("/getall", controllers.GetAllBooks)
	}
}
