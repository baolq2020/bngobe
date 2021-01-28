package models

import (
	"golang.org/x/crypto/bcrypt"
)

func InitDatabase() {

	DB.AutoMigrate(&Book{})
	DB.AutoMigrate(&User{})

	user := DB.Where("username = ?", "admin").First(&User{})

	password := []byte("admin@123")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	if user.RowsAffected == 0 {
		DB.Create(&User{Name: "Superadmin", Email: "admin@admin.com", Username: "admin", Password: string(hashedPassword), IsAdmin: true})
	}

}
