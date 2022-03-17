package models

import "gorm.io/gorm"

type User struct {
	ID       int    `json: "id"`
	Username string `json: "username"`
	Password string `json: "password"`
}

func FindByEmail(db *gorm.DB, username string) (*User, error) {
	var auth User

	result := db.Find(&auth, &User{Username: username})

	return &auth, result.Error
}
