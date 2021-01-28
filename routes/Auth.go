package routes

import (
	"../middleware"
	"../controllers"
	"github.com/gin-gonic/gin"
)

func ApiAuth(route *gin.Engine){

	api := route.Group("/api")
	api.POST("/login", middleware.AuthMiddleware.LoginHandler)
	api.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		api.POST("/check-auth", controllers.CheckAuth)
	}
}