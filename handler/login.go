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
		// Permet la construction du token. Il faudra éventuellement donner l'accesslevel
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		// Permet d'identifier le token, l'interface retournée est ensuite envoyé à "Authorizator"
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			u, _ := su.db.GetUserByEmail(claims[identityKey].(string))
			return u
		},
		// La route /login, permet de donner le jwt token
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			u, _ := su.db.GetUserByEmail(loginVals.Email)
			// Pour afficher l'erreur, remplacer _ par err et décommenter.
			// fmt.Println(err)

			passwordUser := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginVals.Password))
			if loginVals.Email == u.Email && passwordUser == nil {
				return &model.User{
					Email:     u.Email,
					LastName:  "Incomplet",
					FirstName: "Incomplet",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// Donne l'autorisation d'accès ou non. Petite particularité, si l'email du token est "admin@gmail.com" alors on donne le droit tout de même
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.AccessLevel >= 2 || v.Email == "admin@gmail.com" {
				return true
			}

			return false
		},
		// permet de donner le code et le message d'erreur
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

	// Création d'un groupe de route, les routes faites à partir de auth prendront automatiquement "/auth" devant.
	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	// Création des routes protégées à l'intérieur de cette route
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// @path /auth/hello
		auth.GET("/hello", HelloWorld)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
