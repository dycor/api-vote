package postgres

import (
	"fmt"

	"github.com/dycor/api-vote/db"
	"github.com/dycor/api-vote/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresDB struct {
	db *gorm.DB
}

// New is creating a moke to persist data.
func New() db.Persist {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=root sslmode=disable")

	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&model.User{})
	fmt.Println(database.GetErrors())
	postgres := PostgresDB{
		db: database,
	}
	return postgres
}

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
	fmt.Println("Test email")
	var u model.User
	sql.db.Where("email = ?", email).First(&u)
	row := sql.db.Select(&model.User{}).Where("email = ?", email).Row()
	row.Scan(&u)
	fmt.Println(u)
	return nil, nil
}

// GetAllUser is retriving all user form the database.
func (sql PostgresDB) GetAllUser() (map[string]model.User, error) {
	return nil, nil
}
