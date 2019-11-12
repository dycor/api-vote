package handler

import (
	"fmt"
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
	r.GET("/votes", sv.GetAllVoteHandler)
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
	c.JSON(http.StatusNoContent, gin.H{"message": "Record was successfully deleted."})
}

///PutVoteHandler is updating a vote from the given uuid param.
func (sv ServiceVote) PutVoteHandler(ctx *gin.Context) {
	accessLevel := GetAccessLevelJwt(ctx)
	if accessLevel < 1 {
		idUser := GetUUIDJwt(ctx)
		uuid := ctx.Param("uuid")
		if v, error := sv.db.GetVote(uuid); error != nil {
			ctx.JSON(http.StatusInternalServerError, "Vote unknown")
			return
		} else {
			v.UUIDVote = append(v.UUIDVote, idUser)
			if vote, err := sv.db.UpdateVote(uuid, v); err != nil {
				ctx.JSON(http.StatusInternalServerError, "update failed")
				return
			} else {
				fmt.Println("vote", vote)
				ctx.JSON(http.StatusOK, v)
				return
			}
		}
	} else {
		var newVote model.Vote
		if err := ctx.BindJSON(&newVote); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		uuid := ctx.Param("uuid")
		if v, error := sv.db.GetVote(uuid); error != nil {
			ctx.JSON(http.StatusInternalServerError, "Vote unknown")
			return
		} else {
			if newVote.Title != "" {
				v.Title = newVote.Title
			}

			if newVote.Desc != "" {
				v.Desc = newVote.Desc
			}

			if vote, err := sv.db.UpdateVote(uuid, v); err != nil {
				ctx.JSON(http.StatusInternalServerError, "update failed")
				return
			} else {
				fmt.Println("vote", vote)
				ctx.JSON(http.StatusOK, v)
				return
			}
		}
	}

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
	accessLevel := GetAccessLevelJwt(ctx)
	if accessLevel < 1 {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Access level is too slow to access to this route."})
		return
	}
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
