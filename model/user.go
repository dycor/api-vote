package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User is the representation of a client.
type User struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key"`
	AccessLevel int `json:"access_level"`
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth"` // validate:"required"
	//UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt time.Time `json:"deleted_at"`
}

