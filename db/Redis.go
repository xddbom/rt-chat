package db

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func RedisSetup() *redis.Client {
	Client = redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "",
        DB:		  0,  
        Protocol: 2, 
    })

	Ctx := context.Background()

	err := Client.Ping(Ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("Connection error to Redis: %v", err))
	}

	err = Client.Set(Ctx, "Connection", "successful!", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := Client.Get(Ctx, "Connection").Result()
	if err != nil {
		panic(err)
	} 
	fmt.Println("Connection", val)

	return Client
}

func HealthCheck (c *gin.Context, rdb *redis.Client) {
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Redis is not reachable :("})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "message": "Redis is reachable!"})
}