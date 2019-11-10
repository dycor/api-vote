package postgres

import (
	"github.com/dycor/api-vote/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AddVote is adding a vote into the database.
func (sql PostgresDB) AddVote(v *model.Vote) error {
	sql.db.Create(v)
	return nil
}

// DeleteVote is delting a vote from the given UUID.
func (sql PostgresDB) DeleteVote(uuid string) error {
	return nil
}

// UpdateVote is updating a vote from his/here uuid
func (sql PostgresDB) UpdateVote(uuid string, v model.Vote) error {
	return nil
}

// GetVote is getting a vote from his/here uuid.
func (sql PostgresDB) GetVote(uuid string) (*model.Vote, error) {
	var v model.Vote
	err := sql.db.Where(&model.Vote{UUID: uuid}).First(&v).Error
	return &v, err
}

// GetAllVote is retriving all vote form the database.
func (sql PostgresDB) GetAllVote() (map[string]model.Vote, error) {
	return nil, nil
}
