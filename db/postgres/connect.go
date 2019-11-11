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
	database.AutoMigrate(&model.Vote{})

	fmt.Println(database.GetErrors())
	postgres := PostgresDB{
		db: database,
	}
	return postgres
}
