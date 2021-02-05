package routes

import (
	"github.com/gin-gonic/gin"
	"../controllers"
	"../middleware"
)

func ApiRoutes(route *gin.Engine){

	api := route.Group("/api")

	api.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/books", controllers.FindBooks)
			v1.GET("/books/:id", controllers.FindBook)
			v1.POST("/books", controllers.CreateBook)
			v1.PATCH("/books/:id", controllers.UpdateBook)
			v1.DELETE("/books/:id", controllers.DeleteBook)

			v1.POST("/detect/image", controllers.DetectImage)

		}
	}
}