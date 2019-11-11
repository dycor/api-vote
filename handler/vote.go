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
	r.DELETE("/votes/:uuid", sv.DeleteVoteHandler)
	r.PUT("/votes/:uuid", sv.PutVoteHandler)
	r.GET("/votes", sv.GetAllVoteHandler)
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

//DeleteVoteHandler is deleting vote from the given uuid param.
func (sv ServiceVote) DeleteVoteHandler(c *gin.Context) {
	var v model.Vote
	err := sv.db.DeleteVote(c.Param("uuid"), v)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

///PutVoteHandler is updating a vote from the given uuid param.
func (sv ServiceVote) PutVoteHandler(ctx *gin.Context) {
	var newVote model.Vote
	if err := ctx.BindJSON(&newVote); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	uuid := ctx.Param("uuid")
	if err := sv.db.UpdateVote(uuid, newVote); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	v, _ := sv.db.GetVote(uuid)
	ctx.JSON(http.StatusOK, v)
}

//GetAllVoteHandler is retriving all users from the database.
func (sv ServiceVote) GetAllVoteHandler(ctx *gin.Context) {
	var v []model.Vote
	err := sv.db.GetAllVote(&v)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, v)
	}
}

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
