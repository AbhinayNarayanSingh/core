package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ChatController() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Upgrade the HTTP connection to a WebSocket connection
		conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
		if err != nil {
			fmt.Println(err) // Handle error
		}
		fmt.Println(conn)
		// Use the connection
		// ...
	}
}
