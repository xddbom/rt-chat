package main

import (
	"github.com/gin-gonic/gin"

	"github.com/xddbom/rt-chat/routes"
	"github.com/xddbom/rt-chat/db"
)

func main() {
	r := gin.Default()
	
	db.RedisSetup()
	routes.SetupRoutes(r)

	r.Run(":8080")
}