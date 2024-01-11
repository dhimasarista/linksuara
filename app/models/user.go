package models

import "time"

type User struct {
	ID       int
	Username string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User) FindAll() []User {
	var users = []User{
		{
			ID:       1,
			Username: "kotabatam",
			Password: "merdeka2024",
		},
	}

	return users
}
