package moviestudio

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	SaveNewStudio(studio Studio) (Studio, error)
	FindStudioByNumber(studioNumber int) (Studio, error)
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

func (repository *repository) FindStudioByNumber(studioNumber int) (Studio, error) {
	var movieStudio Studio

	err := repository.db.Where("studio_number = ?", studioNumber).Find(&movieStudio).Error
	if (err != nil) {
		return movieStudio, errors.New("failed to find movie studio by number")
	}

	return movieStudio, nil
}
