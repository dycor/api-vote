package handler

import (
	"net/http"
	"regexp"

	"github.com/dycor/api-vote/db"
	"github.com/dycor/api-vote/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

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

// Hash the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// PostUserHandler is creating a new user into the database.
func (su ServiceUser) PostUserHandler(ctx *gin.Context) {

	var u model.User
	if err := ctx.BindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if matched, _ := regexp.Match(`\s|[0-9]+`, []byte(u.FirstName)); matched {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Firstname can't contain numbers and whitespace "})
		return
	}
	if matched, _ := regexp.Match(`\s|[0-9]+`, []byte(u.LastName)); matched {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "LastName can't contain numbers and whitespace "})
		return
	}

	validate = validator.New()
	if err := validate.Struct(u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	u.Password, err = HashPassword(u.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.UUID = uuid.NewV4().String()
	su.db.AddUser(&u)
	ctx.JSON(http.StatusOK, u)

}

func (su ServiceUser) DeleteUserHandler(ctx *gin.Context) {
	accessLevel := GetAccessLevelJwt(ctx)
	if accessLevel < 1 {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Access level is too slow to access to this route."})
		return
	}
	u, err := su.db.GetUser(ctx.Param("uuid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := su.db.DeleteUser(ctx.Param("uuid")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, u)
}
