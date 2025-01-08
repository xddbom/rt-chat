package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func  SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to real time chat based on Redis and Websockets!",
		})
	})


}