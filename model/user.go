package model

import "time"

type User struct {
	Id             string
	Name           string
	Email          string
	PhoneNumber    string
	Password       string
	Address        string
	BirthDate      time.Time
	ProfilePicture string
	CreatedAt      time.Time
	UpdateAt       time.Time
}
