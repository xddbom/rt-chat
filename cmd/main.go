package main

import (
	"github.com/gin-gonic/gin"

	"github.com/xddbom/rt-chat/routes"
	"github.com/xddbom/rt-chat/db"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	rdb := 	db.RedisSetup()	

	routes.SetupRoutes(r, rdb)

	r.Run(":8080")
}