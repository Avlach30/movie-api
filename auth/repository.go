package auth

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	SignUp(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindUserById(id int) (User, error)
	SaveNewUserRole(user User, role string) (error)
	FindUserRoleById(userId int) (UserRole, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) SignUp(user User) (User, error) {
	err := repo.db.Create(&user).Error

	if (err != nil) {
		return user, err
	} else {
		return user, nil
	}
}

func (repo *repository) FindByEmail(email string) (User, error) {
	var user User

	err := repo.db.Where("email = ?", email).Find(&user).Error
	if (err != nil) {
		return user, err
	} else {
		return user, nil
	}
}

func (repo *repository) FindUserById(id int) (User, error) {
	var user User

	err := repo.db.Where("id = ?", id).Find(&user).Error
	if (err != nil) {
		return user, err
	} else {
		return user, nil
	}
}

func (repo *repository) SaveNewUserRole(user User, role string) (error) {
	newRole := UserRole{
		UserId: user.ID,
		Role: role,
	}

	err := repo.db.Create(&newRole).Error
	if err != nil {
		return errors.New("failed to save new user role to database")
	}

	return nil
}

func (repo *repository) FindUserRoleById(userId int) (UserRole, error) {
	var userRole UserRole

	err := repo.db.Where("user_id = ?", userId).Find(&userRole).Error
	if (err != nil) {
		return userRole, errors.New("failed to get user role from database")
	}

	return userRole, nil
}