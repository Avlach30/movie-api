package main

import (
	"movie-api/auth"
	"movie-api/config"
	"movie-api/middleware"
	"movie-api/movie"
	"movie-api/redis"
	moviestudio "movie-api/movie-studio"
	movieschedule "movie-api/movie-schedule"
	movietag "movie-api/movie-tag"
	taskscheduler "movie-api/task-scheduler"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectToDb()

	userRepository := auth.NewRepository(db)
	userService := auth.NewService(userRepository)
	userHandler := auth.NewAuthHandler(userService)

	movieRepository := movie.NewRepository(db)
	movieService := movie.NewService(movieRepository)
	movieHandler := movie.NewMovieHandler(movieService, userService)

	movieStudioRepository := moviestudio.NewRepository(db)
	movieStudioService := moviestudio.NewService(movieStudioRepository)
	movieStudioHandler := moviestudio.NewMovieStudioHandler(movieStudioService, userService)

	movieScheduleRepository := movieschedule.NewRepository(db)
	movieScheduleService := movieschedule.NewService(movieScheduleRepository)
	movieScheduleHandler := movieschedule.NewMovieScheduleHandler(movieScheduleService, movieService, movieStudioService, userService)

	movieTagRepository := movietag.NewRepository(db)
	movieTagService := movietag.NewService(movieTagRepository)
	movieTagHandler := movietag.NewMovieTagHandler(movieTagService)

	redis.RedisConnect()

	router := gin.Default()

	sentry := config.SentryConfig(router)

	router.Use(sentry)

	taskscheduler.NewSchedule()

	//* Configure for accessible static file with first param is router and second param is directory of static file
	router.Static("/avatar", "./avatar")
	router.Static("/movie-poster", "./movie-poster")

	firstVerAPI := router.Group("/api/v1")

	firstVerAPI.POST("/auth/signup-customer", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/admin/signup", userHandler.SignUpHandler)
	firstVerAPI.POST("/auth/login", userHandler.LogInHandler)

	firstVerAPI.GET("/profile", userHandler.GetLoggedUserHandler)

	firstVerAPI.GET("/movies", movieHandler.GetAllMovieWithTags)
	firstVerAPI.GET("/movies/now-playing", middleware.AuthorizationMiddleware(userService), movieScheduleHandler.GetPlayingNowSchedule)

	firstVerAPI.GET("/backoffice/movies",  middleware.AuthorizationMiddleware(userService), movieHandler.GetAllMovieWithTags)
	firstVerAPI.POST("/backoffice/movies", middleware.AuthorizationMiddleware(userService), movieHandler.CreateNewMovieWithTags)

	firstVerAPI.POST("/backoffice/studios", middleware.AuthorizationMiddleware(userService), movieStudioHandler.CreateNewMovieStudio)

	firstVerAPI.POST("/backoffice/schedules", middleware.AuthorizationMiddleware(userService), movieScheduleHandler.CreateNewMovieSchedule)

	firstVerAPI.GET("/backoffice/movies/tags", middleware.AuthorizationMiddleware(userService), movieTagHandler.GetAllTags)

	router.Run()
}