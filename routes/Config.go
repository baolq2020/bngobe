package routes

import (
	"../middleware"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func ConfigRouter() {

	middleware.ConfigAuthMiddleware()

	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}
	// Config CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	// End config CORS

	route := gin.New()
	route.Use(cors.New(config))
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	Handle404Route(route)
	ImportRouter(route)

	if err := http.ListenAndServe(":"+port, route); err != nil {
		log.Fatal(err)
	}

}

func ImportRouter(route *gin.Engine) {
	ApiAuth(route)
	ApiRoutes(route)
}

func Handle404Route(route *gin.Engine) {
	route.NoRoute(middleware.AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
