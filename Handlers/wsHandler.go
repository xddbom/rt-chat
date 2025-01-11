package Handlers

import (
	"log"
	"net/http"
    "time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket" 
    "github.com/xddbom/rt-chat/db"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

var connections = make(map[string]*websocket.Conn)

type WebSocketHandler struct{}

func(h *WebSocketHandler) Handle(c *gin.Context) {
    chatID := c.Query("chatID")
    if chatID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ChatID is required"})
        return
    }

    username := c.GetString("username")
    if username == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
        return
    }

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("Error upgrading connection:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update connection",})
        return
    }
    defer conn.Close()

    connections[username] = conn
    defer delete(connections, username)

	log.Printf("User %s connected", username)

    go db.Subscribe(chatID, func(msg map[string]string) {
        if err := conn.WriteJSON(msg); err != nil {
            log.Println("Error sending message to WebSocket:", err)
        }
    })

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading WebSocket message:", err)
            break
        }

        message := map[string]string{
            "username": username,
            "message":  string(msg),
            "time":     time.Now().Format(time.RFC3339),
        }

        // Сохраняем и публикуем сообщение
        if err := db.SaveMessage(chatID, message); err != nil {
            log.Println("Error saving message to Redis:", err)
        }
        if err := db.Publish(chatID, message); err != nil {
            log.Println("Error publishing message to Redis channel:", err)
        }
    }
}