package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SignUp(input SignUpInput, isAdmin bool) (User, error)
	CheckUserAvailabilityByEmail(input SignUpInput) (bool, error)
	UploadAvatarImage(fileLocation string) (string, error)
	LogIn(input LogInInput) (User, string, error)
	GenerateToken(userId int, email string, isAdmin bool) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) UploadAvatarImage(fileLocation string) (string, error) {
	user := User{}
	user.Avatar = fileLocation

	return fileLocation, nil
}

func (service *service) CheckUserAvailabilityByEmail(input SignUpInput) (bool, error) {
	email := input.Email

	user, err := service.repository.FindByEmail(email)
	if (err != nil) {
		return false, nil
	}

	//* If user exist
	if (user.ID != 0) {
		return true, errors.New("User already exist")
	}

	return false, nil
}

func (service *service) SignUp(input SignUpInput, isAdmin bool) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.IsAdmin = isAdmin

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPw)

	user.Avatar, err = service.UploadAvatarImage(input.Avatar)
	if (err) != nil {
		return user, errors.New("failed to upload image")
	}

	newUser, err := service.repository.SignUp(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (service *service) GenerateToken(userId int, email string, isAdmin bool) (string, error) {

	//* Generate payload token
	claim := jwt.MapClaims{
		"userId": userId,
		"email": email,
		"isAdmin": isAdmin,
	}

	//* Generate token with signing method and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//* Signature token with secret text
	signedToken, err := token.SignedString([]byte("$2a$08$UcyjEygcPA/XaeUp85sQjuOhithx14/Ai3D5lYPixLrMrSQG2NIFy"))
	if (err != nil) {
		return signedToken, err
	}

	return signedToken, nil
}

func (service *service) LogIn(input LogInInput) (User, string, error) {
	email := input.Email
	password := input.Password

	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return user, "", err
	}

	if user.ID == 0 {
		return user, "", errors.New("User with that email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//* If error (compare result is different)
	if err != nil {
		return user, "", errors.New("incorrect Password")
	}

	userToken, err := service.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return user, "", errors.New("failed to generate token")
	}

	return user, userToken, nil
}