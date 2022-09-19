package moviestudio

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	SaveNewStudio(studio Studio) (Studio, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repository *repository) SaveNewStudio(studio Studio) (Studio, error) {
	err := repository.db.Create(&studio).Error
	if err != nil {
		return studio, errors.New("failed to save new movie to database")
	}

	return studio, nil
}
