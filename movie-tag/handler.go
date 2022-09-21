package movietag

import (
	"movie-api/auth"
	"movie-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type movieTagHandler struct {
	service Service
}

func NewMovieTagHandler(service Service) *movieTagHandler {
	return &movieTagHandler{service}
}

func (handler *movieTagHandler) GetAllTags(context *gin.Context) {

	loggedUser := context.MustGet("user").(auth.User)

	if !loggedUser.IsAdmin {
		errorResponse := helper.ApiFailedResponse("Sorry!, only admin can view all movie tags")
		context.JSON(http.StatusForbidden, errorResponse)
		return
	}

	tags, err := handler.service.FetchAllTags()
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiSuccessResponse("Get all tags successfully", tags)
	context.JSON(http.StatusOK, response)
}
