package model

import (
	"time"
)

// User is the representation of a client.
type User struct {
	//gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key"`
	AccessLevel int `json:"access_level"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

