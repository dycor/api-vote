package handler

import (
	"net/http"

	"github.com/dycor/api-vote/db"
	"github.com/dycor/api-vote/model"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// InitVote is creating the endpoints for the entity Vote.
func InitVote(r *gin.Engine, db db.Persist) {
	sv := ServiceVote{
		db: db,
	}
	r.GET("/votes/:uuid", sv.GetVoteHandler)
	//r.DELETE("/users/:uuid", su.DeleteUserHandler)
	//r.PUT("/users/:uuid", su.PutUserHandler)
	//r.GET("/users", su.GetAllUserHandler)
	r.POST("/votes", sv.PostVoteHandler)
}

// ServiceVote is listing all methods on the CURD for REST operations.
type ServiceVote struct {
	db db.Persist
}

// GetVoteHandler is retriving vote from the given uuid param.
func (sv ServiceVote) GetVoteHandler(ctx *gin.Context) {
	v, err := sv.db.GetVote(ctx.Param("uuid"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, v)
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

// PostVoteHandler is creating a new vote into the database.
func (sv ServiceVote) PostVoteHandler(ctx *gin.Context) {
	var v model.Vote
	if err := ctx.BindJSON(&v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate = validator.New()
	if err := validate.Struct(v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v.UUID = uuid.NewV4().String()
	sv.db.AddVote(&v)
	ctx.JSON(http.StatusOK, v)

}
