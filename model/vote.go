package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Vote is the representation of a client.
type Vote struct {
	gorm.Model
	UUID      string         `json:"uuid" gorm:"primary_key"`
	Title     string         `json:"title" validate:"required,min=2"`
	Desc      string         `json:"desc" validate:"required,min=2"`
	StartDate time.Time      `json:"start_date"`
	EndDate   time.Time      `json:"end_date"`
	UUIDVote  pq.StringArray `json:"uuid_votes" gorm:"type:text[]"`
}
