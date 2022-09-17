package movie

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetMoviesWithTags() ([]Movie, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repository *repository) GetMoviesWithTags() ([]Movie, error) {
	var movies []Movie

	err := repository.db.Preload("Tags").Find(&movies).Error
	if (err != nil) {
		return movies, errors.New("failed to get all movies with tags")
	}

	return movies, nil
}