package auth

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Avatar    string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
