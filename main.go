package main

import (
	//"github.com/dycor/api-vote/db/sqlite"
	"github.com/dycor/api-vote/db/postgres"
	"github.com/dycor/api-vote/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler.InitUser(r, postgres.New())
	r.Run(":8080")
}
