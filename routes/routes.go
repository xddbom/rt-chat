package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/xddbom/rt-chat/Handlers"
	"github.com/xddbom/rt-chat/Middlewares"
	"github.com/xddbom/rt-chat/db"
)

func  SetupRoutes(r *gin.Engine, rdb *redis.Client) {
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to real time chat based on Redis and Websockets!",
		})
	})

	r.GET("/health", func(c *gin.Context){
		db.HealthCheck(c, rdb)
	})

	wsHandler := Handlers.WebSocketHandler{}
	r.GET("/ws", middlewares.AuthMiddleware(), wsHandler.Handle)

	r.POST("/login", Handlers.Login)
} 