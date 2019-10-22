package sqlite

import (
	"github.com/dycor/api-vote/db"
	"github.com/dycor/api-vote/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLiteDB struct {
	db *gorm.DB
}

// New is creating a moke to persist data.
func New() db.Persist {
	db, err := gorm.Open("sqlite3", "./api.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	sqlite := SQLiteDB{
		db: db,
	}
	return sqlite
}

// AddUser is adding a user into the database.
func (sql SQLiteDB) AddUser(u *model.User) error {
	sql.db.Create(u)
	return nil
}

// DeleteUser is delting a user from the given UUID.
func (sql SQLiteDB) DeleteUser(uuid string) error {
	return nil
}

// UpdateUser is updating a user from his/here uuid
func (sql SQLiteDB) UpdateUser(uuid string, u model.User) error {
	return nil
}

// GetUser is getting a user from his/here uuid.
func (sql SQLiteDB) GetUser(uuid string) (*model.User, error) {
	return nil, nil
}

// GetAllUser is retriving all user form the database.
func (sql SQLiteDB) GetAllUser() (map[string]model.User, error) {
	return nil, nil
}
