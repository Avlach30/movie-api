package auth

import "gorm.io/gorm"

type Repository interface {
	SignUp(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindUserById(id int) (User, error)
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