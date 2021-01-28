package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAuth(c *gin.Context) {

	CurrentUser := models.GetCurrentUser(c)

	c.JSON(http.StatusOK, gin.H{"username": CurrentUser.Username, "email": CurrentUser.Email, "name": CurrentUser.Name, "is_admin": CurrentUser.IsAdmin})

}
