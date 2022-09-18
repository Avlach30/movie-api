package movie

import (
	"movie-api/auth"
	"movie-api/helper"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type movieHandler struct {
	service Service
}

func NewMovieHandler(service Service) *movieHandler {
	return &movieHandler{service}
}

func (handler *movieHandler) GetAllMovieWithTags(context *gin.Context) {

	if (context.Request.URL.String() == "/api/v1/backoffice/movies") {
		loggedUser := context.MustGet("user").(auth.User)

		//* If logged user not found (invalid token / authorization header)
		if (loggedUser.ID == 0) {
			errorResponse := helper.ApiFailedResponse("Unauthorized")
			context.JSON(http.StatusUnauthorized, errorResponse)
			return
		}
	}

	movies, err := handler.service.FetchAllMovieWithTags()
	if (err != nil) {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiSuccessResponse("Get all movie with tags successfully", FormatResponseGetMovies(movies))
	context.JSON(http.StatusOK, response)
}

func (handler *movieHandler) CreateNewMovieWithTags(context *gin.Context) {
	loggedUser := context.MustGet("user").(auth.User)

	if (!loggedUser.IsAdmin) {
		errorResponse := helper.ApiFailedResponse("Sorry!, only admin can create new movie")
		context.JSON(http.StatusForbidden, errorResponse)
		return
	}

	input := CreateNewMovieInput{
		Title: context.PostForm("title"),
		Overview: context.PostForm("overview"),
		PlayUntil: context.PostForm("play_until"),
	}

	err := context.ShouldBind(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMsg := gin.H{"errors": errors}

		errorResponse := helper.ApiFailedResponse(errorMsg)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}


	input.Tags = context.PostFormArray("tags")
	

	imagePoster, err := context.FormFile("poster")
	if imagePoster == nil {
		errorResponse := helper.ApiFailedResponse("Sorry, you must upload an poster image")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	extension := filepath.Ext(imagePoster.Filename)

	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
		errorResponse := helper.ApiFailedResponse("Sorry, only image file type can be uploaded")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if imagePoster.Size > 1572864 {
		errorResponse := helper.ApiFailedResponse("Sorry, image uploaded is more than limit (1,5 mb)")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	newFileName := uuid.New().String() + extension

	imagePath := "movie-poster/" + newFileName

	err = context.SaveUploadedFile(imagePoster, imagePath)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	input.Poster = imagePath


	newMovie, err := handler.service.SaveNewMovieWithTags(input)
	if (err != nil) {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiSuccessResponse("Create new movie with tags successfully", FormatResponseCreateNewMovie(newMovie))
	context.JSON(http.StatusCreated, response)
}