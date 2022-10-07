package movieschedule

import (
	"movie-api/movie"
	moviestudio "movie-api/movie-studio"
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
	Movie     movie.Movie
	Studio    moviestudio.Studio
}
