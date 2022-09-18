package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB) {
	dsn := "root:rootPassword!@tcp(127.0.0.1:3306)/movie-api-onboarding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to database!")

	return db
}