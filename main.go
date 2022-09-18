package main

import (
	"log"
	"movie-api/auth"
	"movie-api/movie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:rootPassword!@tcp(127.0.0.1:3306)/movie-api-onboarding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to database!")

	userRepository := auth.NewRepository(db)
	userService := auth.NewService(userRepository)
	userHandler := auth.NewAuthHandler(userService)

	movieRepository := movie.NewRepository(db)
	movieService := movie.NewService(movieRepository)
	movieHandler := movie.NewMovieHandler(movieService)

	router := gin.Default()
	firstVerAPI := router.Group("/api/v1")

	firstVerAPI.POST("/auth/signup-customer", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/admin/signup", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/login", userHandler.LogInHandler)

	firstVerAPI.GET("/movies", movieHandler.GetAllMovieWithTags)

	firstVerAPI.GET("/backoffice/movies", movieHandler.GetAllMovieWithTags)
	firstVerAPI.POST("/backoffice/movies", movieHandler.CreateNewMovieWithTags)

	router.Run()
}