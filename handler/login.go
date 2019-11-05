package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dycor/api-vote/db"
	"github.com/dycor/api-vote/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func GetToken(c *gin.Context) string {
	token, exists := c.Get("JWT_TOKEN")
	if !exists {
		return ""
	}
	return token.(string)
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"text":  "Hello World.",
		"token": GetToken(c),
	})
	token := GetToken(c)
	fmt.Println("TOKEN ------>")
	fmt.Println(token)
}

// @Path = /auth/hello
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"userID": "ttetet",
		"email":  "tet",
		"text":   "Hello World.",
	})
}

// InitLogin is creating jwt Token for users
func InitLogin(r *gin.Engine, port string, db db.Persist) {
	su := ServiceUser{
		db: db,
	}
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			// pas sur de ca
			u, _ := su.db.GetUserByEmail(claims[identityKey].(string))
			return u
		},
		// La route /login
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			u, _ := su.db.GetUserByEmail(loginVals.Email)
			// Pour afficher l'erreur, remplacer _ par err et décommenter
			// fmt.Println(err)

			passwordUser := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginVals.Password))
			if loginVals.Email == u.Email && passwordUser == nil {
				//Pas sur de ce que ca doit retourner hihihi
				return &model.User{
					Email:     u.Email,
					LastName:  "Test",
					FirstName: "Test",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// Cette fonction permet de passer outre le jwt == Si email == l'adresse indiqué, alors il a automatiquement le droit
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.Email == "admin@gmail.com" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	//auth.GET("/hello2", helloHandler)
	{
		// @path /auth/hello
		auth.GET("/hello", HelloWorld)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
