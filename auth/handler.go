package auth

import (
	"movie-api/helper"
	"movie-api/redis"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type authHandler struct {
	authService Service
}

func NewAuthHandler(authService Service) *authHandler {
	return &authHandler{authService}
}

func (handler *authHandler) SignUpHandler(context *gin.Context) {
	input := SignUpInput{
		Name:        context.PostForm("name"),
		Email:       context.PostForm("email"),
		PhoneNumber: context.PostForm("phone_number"),
		Password:    context.PostForm("password"),
	}

	err := context.ShouldBind(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorResponse := helper.ApiFailedResponse(errors)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	imageAvatar, err := context.FormFile("avatar")
	if imageAvatar == nil {
		errorResponse := helper.ApiFailedResponse("Sorry, you must upload an avatar image")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	extension := filepath.Ext(imageAvatar.Filename)

	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
		errorResponse := helper.ApiFailedResponse("Sorry, only image file type can be uploaded")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if imageAvatar.Size > 1572864 {
		errorResponse := helper.ApiFailedResponse("Sorry, image uploaded is more than limit (1,5 mb)")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	newFileName := uuid.New().String() + extension

	imagePath := "avatar/" + newFileName

	err = context.SaveUploadedFile(imageAvatar, imagePath)
	if err != nil {
		errorResponse := helper.ApiFailedResponse("failed to save uploaded file")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	input.Avatar = imagePath

	isUserExist, err := handler.authService.CheckUserAvailabilityByEmail(input)
	if isUserExist {
		errorResponse := helper.ApiFailedResponse(err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	endpoint := context.Request.URL.String()

	if endpoint == "/api/v1/auth/signup-customer" {
		user, err := handler.authService.SignUp(input, false)
		if err != nil {
			errorResponse := helper.ApiFailedResponse(err.Error())
			context.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		responseFormatter := FormatUserSignupResponse(user, imagePath)
		successResponse := helper.ApiSuccessResponse("Sign up for customer successfully", responseFormatter)

		context.JSON(http.StatusCreated, successResponse)
	} else {
		user, err := handler.authService.SignUp(input, true)
		if err != nil {
			errorResponse := helper.ApiFailedResponse(err.Error())
			context.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		responseFormatter := FormatUserSignupResponse(user, imagePath)
		successResponse := helper.ApiSuccessResponse("Sign up for admin successfully", responseFormatter)

		context.JSON(http.StatusCreated, successResponse)
	}
}

func (handler *authHandler) LogInHandler(context *gin.Context) {
	var input LogInInput //* Mendefinisikan variabel dengan type LogInInput struct

	err := context.ShouldBindJSON(&input) //*Konversi JSON input ke struct
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorResponse := helper.ApiFailedResponse(errors)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	user, token, err := handler.authService.LogIn(input)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusUnauthorized, errorResponse)
		return
	}

	responseFormatter := FormatUserLoginResponse(user, token)
	successResponse := helper.ApiSuccessResponse("Log in successfully", responseFormatter)

	context.JSON(http.StatusOK, successResponse)
}

func (handler *authHandler) GetLoggedUserHandler(context *gin.Context) {
	//* Calling service for get value from redis key
	loggedUserId, err := redis.GetRedisFromKey("userId")
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusUnauthorized, errorResponse)
		return
	}

	user, err := handler.authService.FindUserById(loggedUserId)
	if err != nil {
		errorResponse := helper.ApiFailedResponse(err.Error())

		context.JSON(http.StatusNotFound, errorResponse)
		return
	}

	responseFormatter := FormatGetLoggedUserResponse(user)
	successResponse := helper.ApiSuccessResponse("Get logged user successfully", responseFormatter)

	context.JSON(http.StatusOK, successResponse)
}
