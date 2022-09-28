package auth

type SignUpInput struct {
	Name        string `form:"name" binding:"required"`
	Email       string `form:"email" binding:"required,email"`
	PhoneNumber string `form:"phone_number" binding:"required,e164"`
	Password    string `form:"password" binding:"required,min=10,alphanum"`
	Avatar      string
}

type LogInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
