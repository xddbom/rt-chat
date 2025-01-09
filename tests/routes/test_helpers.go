package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/xddbom/rt-chat/routes"
	"github.com/xddbom/rt-chat/db"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	rdb := db.RedisSetup()
	routes.SetupRoutes(r, rdb)
	return r
}