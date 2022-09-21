package movieschedule

import (
	"time"
)

type MovieSchedule struct {
	ID        int
	MovieId   int
	StudioId  int
	StartTime string
	EndTime   string
	Price     int
	Date      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}