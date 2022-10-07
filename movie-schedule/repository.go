package movieschedule

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository interface{
	SaveNewSchedule(MovieSchedule) (MovieSchedule, error)
	GetPlayingNowSchedule() ([]MovieSchedule, error)
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

func (repository *repository) GetPlayingNowSchedule() ([]MovieSchedule, error) {
	var schedules []MovieSchedule

	currentDate := time.Now().Format("2006-01-02")

	err := repository.db.Preload("Movie").Preload("Studio").Where("date LIKE ?", currentDate + "%").Find(&schedules).Error
	if (err != nil) {
		return schedules, errors.New("failed to get playing now shcedule from database")
	}


	return schedules, nil
}
