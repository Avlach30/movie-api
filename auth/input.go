package auth

type SignUpInput struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=10,alphanum,uppercase,lowercase"`
	Avatar   string `form:"avatar" binding:"required"`
}
