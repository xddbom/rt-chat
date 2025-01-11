package db

import (
	"context"
	"fmt"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func RedisSetup() *redis.Client {
	Client = redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "",
        DB:		  0,  
        Protocol: 2, 
    })

	Ctx = context.Background()

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

	dbSize, err := rdb.DBSize(ctx).Result()
    if err != nil {
        c.JSON(500, gin.H{"status": "error", "message": "Failed to fetch Redis DB size"})
        return
    }

    c.JSON(200, gin.H{
        "status":   "ok",
        "message":  "Redis is reachable!",
        "db_size":  dbSize,
    })
}


func SaveMessage(chatID string, message map[string]string) error {
    key := fmt.Sprintf("chat:%s:messages", chatID)
    messageJSON, err := json.Marshal(message)
    if err != nil {
        return err
    }

	if err := Client.LPush(Ctx, key, messageJSON).Err(); err != nil {
        return err
    }

    return Client.LTrim(Ctx, key, 0, 99).Err()
}



func GetChatHistroy(chatID string, limit int64) ([]map[string]string, error) {
	key := fmt.Sprintf("chat:%s:messages", chatID)
	messages, err := Client.LRange(Ctx, key, 0, limit-1).Result()
	if err != nil {
		return nil, err
	}

	var history []map[string]string
	for _, msg := range messages {
		var message map[string]string 
		if err := json.Unmarshal([]byte(msg), &message); err == nil {
			history = append(history, message)
		}
	}

	return history, nil
}


func Publish(chatID string, message map[string]string) error {
    key := fmt.Sprintf("chat:%s:channel", chatID)
    messageJSON, err := json.Marshal(message)
    if err != nil {
        return err
    }

    return Client.Publish(Ctx, key, messageJSON).Err()
}


func Subscribe(chatID string, messageHandler func(msg map[string]string)) {
    key := fmt.Sprintf("chat:%s:channel", chatID)
    pubsub := Client.Subscribe(Ctx, key)
    defer pubsub.Close()

    ch := pubsub.Channel()
    for msg := range ch {
        var message map[string]string
        if err := json.Unmarshal([]byte(msg.Payload), &message); err == nil {
            messageHandler(message)
        }
    }
}