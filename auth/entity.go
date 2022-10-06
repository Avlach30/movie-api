package auth

import "time"

type User struct {
	ID          int
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	Avatar      string
	IsAdmin     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type UserRole struct {
	ID     int
	UserId int
	Role   string
}
