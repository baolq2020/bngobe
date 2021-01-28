package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

const CurrentUserKey string = "current-user"

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email  string `json:"email" gorm:"size:255;unique;not null"`
	Username  string `json:"username" gorm:"size:255;unique;not null"`
	Password string `json:"password"`
	IsAdmin bool `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetCurrentUser(c *gin.Context) User  {

	objUser, _ := c.MustGet(CurrentUserKey).(User)

	return objUser

}
