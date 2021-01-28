package middleware

import (
	"../models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func ConfigAuthMiddleware()  {

	var identityKey = "id"

	var secretKey = os.Getenv("SECRET_KEY")
	if secretKey == ""{
		secretKey = "secret key"
	}

	ginJWT, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "production",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			var User models.User
			query := models.DB.Where("username = ? OR email = ?", loginVals.Username, loginVals.Username).First(&User)

			if query.RowsAffected > 0{
				err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password))
				if err == nil{
					return &models.User{
						Username:  userID,
					}, nil
				}
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {

			token, ok := data.(*models.User)

			if ok {
				var User models.User
				query := models.DB.Where("username = ? OR email = ?", token.Username, token.Username).First(&User)
				if query.RowsAffected == 0{
					return false
				}
				c.Set(models.CurrentUserKey, User)
				c.Next()
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	AuthMiddleware = ginJWT

	errInit := AuthMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

}

