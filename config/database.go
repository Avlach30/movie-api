package config

import (
	"fmt"
	"log"
	"movie-api/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB) {

	get := helper.GetEnvValue
	DB_USER := get("DB_USER")
	DB_PASSWORD := get("DB_PASSWORD")
	DB_HOST := get("DB_HOST")
	DB_PORT := get("DB_PORT")
	DB_DATABASE := get("DB_DATABASE")
	
	dsn :=  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to database!")

	return db
}