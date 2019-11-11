package db

import (
	"errors"

	"github.com/dycor/api-vote/model"
)

// Persist is a interface contract that define CRUD into the database.
type Persist interface {
	AddUser(u *model.User) error
	DeleteUser(uuid string) error
	UpdateUser(uuid string, u *model.User) (*model.User, error)
	GetUser(uuid string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetAllUser() (map[string]model.User, error)
}

// ErrUserNotFound is used if a user is not found into the DB.
var ErrUserNotFound = errors.New("db: user not found")
