package movietag

import "time"

type Tag struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type MovieTag struct {
	ID        int
	MovieId   int
	TagId     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
