package main

import (
	//"github.com/dycor/api-vote/db/sqlite"
	"github.com/dycor/api-vote/db/postgres"
	"github.com/dycor/api-vote/handler"

	"github.com/gin-gonic/gin"
)

var port string

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	port := "8080"
	if port == "" {
		port = "8080"
	}

	handler.InitUser(r, postgres.New())
	handler.InitVote(r, postgres.New())
	handler.InitLogin(r, port, postgres.New())
	r.Run(":" + port)
}
