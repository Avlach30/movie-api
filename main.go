package main

import (
	"movie-api/auth"
	"movie-api/config"
	"movie-api/middleware"
	"movie-api/movie"
	moviestudio "movie-api/movie-studio"
	movieschedule "movie-api/movie-schedule"
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

	movieStudioRepository := moviestudio.NewRepository(db)
	movieStudioService := moviestudio.NewService(movieStudioRepository)
	movieStudioHandler := moviestudio.NewMovieStudioHandler(movieStudioService)

	movieScheduleRepository := movieschedule.NewRepository(db)
	movieScheduleService := movieschedule.NewService(movieScheduleRepository)
	movieScheduleHandler := movieschedule.NewMovieScheduleHandler(movieScheduleService, movieService, movieStudioService)

	router := gin.Default()
	firstVerAPI := router.Group("/api/v1")

	firstVerAPI.POST("/auth/signup-customer", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/admin/signup", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/login", userHandler.LogInHandler)

	firstVerAPI.GET("/movies", movieHandler.GetAllMovieWithTags)

	firstVerAPI.GET("/backoffice/movies",  middleware.AuthorizationMiddleware(userService), movieHandler.GetAllMovieWithTags)
	firstVerAPI.POST("/backoffice/movies", middleware.AuthorizationMiddleware(userService), movieHandler.CreateNewMovieWithTags)

	firstVerAPI.POST("/backoffice/studios", middleware.AuthorizationMiddleware(userService), movieStudioHandler.CreateNewMovieStudio)

	firstVerAPI.POST("/backoffice/schedules", middleware.AuthorizationMiddleware(userService), movieScheduleHandler.CreateNewMovieSchedule)

	router.Run()
}