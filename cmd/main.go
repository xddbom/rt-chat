package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xddbom/rt-chat/routes"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}