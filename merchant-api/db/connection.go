package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:majoo@tcp(merchant-api-db:3306)/majootest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot Connect")
	}

	return db
}
