package moviestudio

import "time"

type Studio struct {
	ID           int
	StudioNumber int
	SeatCapacity int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
