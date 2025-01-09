package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/xddbom/rt-chat/routes"
	"github.com/xddbom/rt-chat/db"
)

func main() {
	r := gin.Default()
	rdb := 	db.RedisSetup()	

	routes.SetupRoutes(r, rdb)

	fmt.Println("test ing workflow")

	r.Run(":8080")
}