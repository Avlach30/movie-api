package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SignUp(input SignUpInput, isAdmin bool) (User, error)
	CheckUserAvailabilityByEmail(input SignUpInput) (bool, error)
	UploadAvatarImage(fileLocation string) (string, error)
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