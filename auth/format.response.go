package auth

type SignUpResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type LoginResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Token  string `json:"_token"`
}

func FormatUserSignupResponse(user User, imageLocation string) SignUpResponse {
	format := SignUpResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: imageLocation,
	}

	return format
}

func FormatUserLoginResponse(user User, token string) LoginResponse {
	format := LoginResponse{
		Name: user.Name,
		Email: user.Email,
		Avatar: user.Avatar,
		Token: token,
	}

	return format
}