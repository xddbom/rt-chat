package tests

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/xddbom/rt-chat/Handlers"
)

func TestWsHandler(t *testing.T) {
	serverReady := make(chan struct{})

	go func() {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		wsHandler := Handlers.WebSocketHandler{}
		r.GET("/ws", wsHandler.Handle)
		
		go func() {
			if err := r.Run(":8080"); err != nil {
				t.Errorf("Server failed to start: %v", err)
			}
		}()

		close(serverReady)
	}()

	<-serverReady

	time.Sleep(500 * time.Millisecond)

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws?username=testUser", nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte("Test"))
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	assert.Equal(t, "Test", string(msg))
}