package movieschedule

import (
	"movie-api/auth"
	"movie-api/helper"
	"movie-api/movie"
	moviestudio "movie-api/movie-studio"
	"net/http"

	"github.com/gin-gonic/gin"
)

type movieScheduleHandler struct {
	movieScheduleService Service
	movieService movie.Service
	movieStudioService moviestudio.Service
	authService auth.Service
}

func NewMovieScheduleHandler(movieScheduleService Service, movieService movie.Service, movieStudioService moviestudio.Service, authService auth.Service) *movieScheduleHandler {
	return &movieScheduleHandler{movieScheduleService, movieService, movieStudioService, authService}
}

func (handler *movieScheduleHandler) CreateNewMovieSchedule(context *gin.Context) {
	loggedUser := context.MustGet("user").(auth.User)

	roleUser, err := handler.authService.FindRoleUserById(loggedUser.ID)
	if (err != nil) {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if (roleUser.Role != "full-control") {
		errorResponse := helper.ApiFailedResponse("Sorry!, only admin can create new schedule for playing movie")
		context.JSON(http.StatusForbidden, errorResponse)
		return
	}

	var input CreateNewScheduleInput

	err = context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorResponse := helper.ApiFailedResponse(errors)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	movie, err := handler.movieService.FetchMovieById(input.MovieId)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	//* If movie not found
	if (movie.ID == 0) {
		errorResponse := helper.ApiFailedResponse("Sorry! movie not found")

		context.JSON(http.StatusNotFound, errorResponse)
		return
	}

	studio, err := handler.movieStudioService.FindStudioByID(input.StudioId)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	//* If studio not found
	if (studio.ID == 0) {
		errorResponse := helper.ApiFailedResponse("Sorry! studio not found")

		context.JSON(http.StatusNotFound, errorResponse)
		return
	}


	movieSchedule, err := handler.movieScheduleService.SaveNewMovieSchedule(input)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}
	
	successResponse := helper.ApiSuccessResponse("Create new schedule for playing movie successfully", FormatCreateNewScheduleResponse(movieSchedule, movie, studio))

	context.JSON(http.StatusCreated, successResponse)
}

func (handler *movieScheduleHandler) GetPlayingNowSchedule(context *gin.Context) {
	loggedUser := context.MustGet("user").(auth.User)

	//* If logged user not found (invalid token / authorization header)
	if (loggedUser.ID == 0) {
		errorResponse := helper.ApiFailedResponse("Unauthorized")
		context.JSON(http.StatusUnauthorized, errorResponse)
		return
	}

	playingNowSchedules, err := handler.movieScheduleService.FetchPlayingNowSchedule()
	if (err != nil) {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	response := helper.ApiSuccessResponse("Get all playing now schedule successfully", FormatGetPlayingNowSchedulesResponse(playingNowSchedules))
	context.JSON(http.StatusOK, response)
}