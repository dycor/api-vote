package handler

import (
	"net/http"

	"github.com/dycor/api-vote/db"
	"github.com/dycor/api-vote/model"

	"github.com/gin-gonic/gin"
)

// InitUser is creating the endpoints for the entity User.
func InitUser(r *gin.Engine, db db.Persist) {
	su := ServiceUser{
		db: db,
	}
	r.GET("/users/:uuid", su.GetUserHandler)
	//r.DELETE("/users/:uuid", su.DeleteUserHandler)
	//r.PUT("/users/:uuid", su.PutUserHandler)
	//r.GET("/users", su.GetAllUserHandler)
	r.POST("/users", su.PostUserHandler)
}

// ServiceUser is listing all methods on the CURD for REST operations.
type ServiceUser struct {
	db db.Persist
}

// GetUserHandler is retriving user from the given uuid param.
func (su ServiceUser) GetUserHandler(ctx *gin.Context) {
	u, err := su.db.GetUser(ctx.Param("uuid"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, u)
}
//
//// DeleteUserHandler is deleting a user from the given uuid param.
//func (su ServiceUser) DeleteUserHandler(ctx *gin.Context) {
//	err := su.db.DeleteUser(ctx.Param("uuid"))
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, nil)
//		return
//	}
//	ctx.JSON(http.StatusOK, nil)
//}
//
//// PutUserHandler is updating a user from the given uuid param.
//func (su ServiceUser) PutUserHandler(ctx *gin.Context) {
//
//	var newUser model.User
//	if err := ctx.BindJSON(&newUser); err != nil {
//		ctx.JSON(http.StatusBadRequest, nil)
//		return
//	}
//	uuid := ctx.Param("uuid")
//	if err := su.db.UpdateUser(uuid, newUser); err != nil {
//		ctx.JSON(http.StatusInternalServerError, nil)
//		return
//	}
//	u, _ := su.db.GetUser(uuid)
//	ctx.JSON(http.StatusOK, u)
//}

// GetAllUserHandler is retriving all users from the database.
//func (su ServiceUser) GetAllUserHandler(ctx *gin.Context) {
//	us, _ := su.db.GetAllUser()
//	ctx.JSON(http.StatusOK, us)
//}

// PostUserHandler is creating a new user into the database.
func (su ServiceUser) PostUserHandler(ctx *gin.Context) {
	var u model.User
	if err := ctx.BindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	su.db.AddUser(&u)
	ctx.JSON(http.StatusOK, u)
}
