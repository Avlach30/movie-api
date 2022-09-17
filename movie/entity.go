package movie

import (
	movietag "movie-api/movie-tag"
	"time"
)

type Movie struct {
	ID        int
	Title     string
	Overview  string
	Poster    string
	PlayUntil string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Tags      []movietag.Tag `gorm:"many2many:movie_tags"`
}
