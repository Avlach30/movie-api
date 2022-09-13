package auth

type SignUpResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Avatar	   string `json:"avatar"`
}

func FormatUserSignupResponse(user User, imageLocation string) SignUpResponse {
	format := SignUpResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Avatar:		imageLocation,
	}

	return format
}