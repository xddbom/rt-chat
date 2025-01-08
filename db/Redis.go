package db

import (
	"context"
	"fmt"

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