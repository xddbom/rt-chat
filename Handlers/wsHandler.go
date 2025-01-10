package Handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

var connections = make(map[string]*websocket.Conn)

type WebSocketHandler struct{}

func(h *WebSocketHandler) Handle(c *gin.Context) {
    username := c.DefaultQuery("username", "")                      // query || c.Get (?)
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
    defer func() {
        delete(connections, username)
        conn.Close()
    }()


	log.Printf("User %s connected", username)

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }

        log.Printf("[%s]: %s", username, msg)

        for _, c := range connections {
            if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
                log.Println("Error sending message:", err)
            }
        }
    }
}