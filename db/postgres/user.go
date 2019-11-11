package postgres

import (
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
	return nil
}

// UpdateUser is updating a user from his/here uuid
func (sql PostgresDB) UpdateUser(uuid string, u model.User) error {
	return nil
}

// GetUser is getting a user from his/here uuid.
func (sql PostgresDB) GetUser(uuid string) (*model.User, error) {
	return nil, nil
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
