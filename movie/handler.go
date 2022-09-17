package movie

import (
	"movie-api/helper"
	"net/http"
	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	service Service
}

func NewMovieHandler(service Service) *movieHandler {
	return &movieHandler{service}
}

func (handler *movieHandler) GetAllMovieWithTags(context *gin.Context) {
	movies, err := handler.service.FetchAllMovieWithTags()
	if (err != nil) {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiSuccessResponse("Get all movie with tags successfully", FormatResponseGetMovies(movies))
	context.JSON(http.StatusOK, response)
}