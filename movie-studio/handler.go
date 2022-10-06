package moviestudio

import (
	"movie-api/auth"
	"movie-api/helper"
	"net/http"
	"github.com/gin-gonic/gin"
)

type movieStudioHandler struct {
	movieStudioService Service
	authService auth.Service
}

func NewMovieStudioHandler(movieStudioService Service, authService auth.Service) *movieStudioHandler {
	return &movieStudioHandler{movieStudioService, authService}
}

func (handler *movieStudioHandler) CreateNewMovieStudio(context *gin.Context) {
	loggedUser := context.MustGet("user").(auth.User)

	roleUser, err := handler.authService.FindRoleUserById(loggedUser.ID)
	if (err != nil) {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if (roleUser.Role != "full-control") {
		errorResponse := helper.ApiFailedResponse("Sorry!, only admin can create new movie studio")
		context.JSON(http.StatusForbidden, errorResponse)
		return
	}

	var input CreateNewStudioInput

	err = context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorResponse := helper.ApiFailedResponse(errors)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	studio, err := handler.movieStudioService.FindStudioByNumber(input.StudioNumber)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	//* If movie studio already exist
	if (studio.ID != 0) {
		errorResponse := helper.ApiFailedResponse("Sorry! movie studio already exist")

		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	movieStudio, err := handler.movieStudioService.SaveNewStudio(input)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusUnauthorized, errorResponse)
		return
	}
	
	successResponse := helper.ApiSuccessResponse("Create new movie studio successfully", FormatCreateNewStudioResponse(movieStudio))

	context.JSON(http.StatusCreated, successResponse)
}