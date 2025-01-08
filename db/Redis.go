package db

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RedisSetup() *redis.Client {
	client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "",
        DB:		  0,  
        Protocol: 2, 
    })

	Ctx := context.Background()

	err := client.Ping(Ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("Connection error to Redis: %v", err))
	}

	err = client.Set(Ctx, "Connection", "successful!", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(Ctx, "Connection").Result()
	if err != nil {
		panic(err)
	} 
	fmt.Println("Connection", val)

	return client
}

func HealthCheckc (c *gin.Context, rdb *redis.Client) {
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Redis is not reachable :("})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "message": "Redis is reachable!"})
}