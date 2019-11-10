package moke

import (
	"github.com/dycor/api-vote/model"
	uuid "github.com/satori/go.uuid"
)

// ListVote is a moke of a DB
type ListVote map[string]model.Vote

// AddUser is adding a new user into the database.
func (lu ListVote) AddVote(v *model.Vote) error {
	v.UUID = uuid.NewV4().String()
	lu[v.UUID] = *v
	return nil
}

//// DeleteUser is deleting a User into the database.
//func (lu ListUser) DeleteUser(uuid string) error {
//	if _, ok := lu[uuid]; !ok {
//		return db.ErrUserNotFound
//	}
//	delete(lu, uuid)
//	return nil
//}
//
//// UpdateUser user is updating a User form the given uuid and user.
//func (lu ListUser) UpdateUser(uuid string, newUser model.User) error {
//	u, ok := lu[uuid]
//	if !ok {
//		return db.ErrUserNotFound
//	}
//	newUser.UUID = u.FirstName
//	u.FirstName = newUser.FirstName
//	u.LastName = newUser.LastName
//	lu[uuid] = u
//	return nil
//}
//
//// GetUser is retriving from the database the given uuid user.
//func (lu ListUser) GetUser(uuid string) (*model.User, error) {
//	u := lu[uuid]
//	return &u, nil
//}
//
//// GetAllUser is getting all users from the database.
//func (lu ListUser) GetAllUser() (map[string]model.User, error) {
//	return lu, nil
//}
