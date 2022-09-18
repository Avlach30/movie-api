package main

import (
	"movie-api/auth"
	"movie-api/config"
	"movie-api/movie"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectToDb()

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