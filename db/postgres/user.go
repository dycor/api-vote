package postgres

import (
	"time"

	"github.com/dycor/api-vote/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AddUser is adding a user into the database.
func (sql PostgresDB) AddUser(u *model.User) error {
	sql.db.Create(u)
	return nil
}

// DeleteUser is delting a user from the given UUID.
func (sql PostgresDB) DeleteUser(uuid string) error {
	t := time.Now()
	err := sql.db.Exec("UPDATE users SET deleted_at=$1 WHERE uuid=$2", t, uuid)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateUser is updating a user from his/here uuid
func (sql PostgresDB) UpdateUser(uuid string, u model.User) error {
	return nil
}

// GetUser is getting a user from his/here uuid.
func (sql PostgresDB) GetUser(uuid string) (*model.User, error) {
	var u model.User
	err := sql.db.Where(&model.User{UUID: uuid}).First(&u).Error
	return &u, err
}

// GetUser is getting a user from his/here email.
func (sql PostgresDB) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	err := sql.db.Where(&model.User{Email: email}).First(&u).Error
	return &u, err
}

// GetAllUser is retriving all user form the database.
func (sql PostgresDB) GetAllUser() (map[string]model.User, error) {
	return nil, nil
}
