package auth

import "strconv"

type SignUpResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	Avatar      string `json:"avatar"`
}

type LoginResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Token  string `json:"_token"`
}

func FormatUserSignupResponse(user User, imageLocation string) SignUpResponse {

	IntPhoneNumber, _ := strconv.Atoi(user.PhoneNumber)

	format := SignUpResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: IntPhoneNumber,
		Avatar:      "/" + imageLocation,
	}

	return format
}

func FormatUserLoginResponse(user User, token string) LoginResponse {
	format := LoginResponse{
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
		Token:  token,
	}

	return format
}

func FormatGetLoggedUserResponse(user User) SignUpResponse {
	IntPhoneNumber, _ := strconv.Atoi(user.PhoneNumber)

	format := SignUpResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: IntPhoneNumber,
		Avatar:      user.Avatar,
	}

	return format
}
