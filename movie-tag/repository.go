package movietag

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	FindAllTags() ([]Tag, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repository *repository) FindAllTags() ([]Tag, error) {
	var tags []Tag

	err := repository.db.Find(&tags).Error
	if (err != nil) {
		return tags, errors.New("failed to get all movie tags")
	}

	return tags, nil
}