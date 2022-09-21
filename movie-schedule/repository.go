package movieschedule

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface{
	SaveNewSchedule(MovieSchedule) (MovieSchedule, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repository *repository) SaveNewSchedule(schedule MovieSchedule) (MovieSchedule, error) {
	err := repository.db.Create(&schedule).Error
	if err != nil {
		return schedule, errors.New("failed to save new movie playing schedule to database")
	}

	return schedule, nil
}
