package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	PUBLISH        = "PUBLISH"
	SUBSCRIBE      = "SUBSCRIBE"
	UNSUBSCRIBE    = "UNSUBSCRIBE"
	FETCH_MESSAGE  = "FETCH_MESSAGE"
	ONLINE_STATUS  = "ONLINE_STATUS"
	OFFLINE_STATUS = "OFFLINE_STATUS"
)

type Client struct {
	Id         string
	Connection *websocket.Conn
}

type Payload struct {
	Action   string `json:"action"`
	Group_Id string `json:"group_id"`
	Message  string `json:"message"`
}

type Subscription struct {
	Topic  string
	Client *Client
}

func autoId() string {
	return uuid.New().String()
}

var Clients []Client

func ReceiveMessageHandler(conn websocket.Conn, messageType int, payload []byte) {

	msg := Payload{}

	fmt.Printf("%s\n", payload)
	fmt.Printf("%T\n", payload)

	if err := conn.WriteMessage(messageType, payload); err != nil {
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal([]byte(payload), &msg); err != nil {
		fmt.Println(err)
		if err := conn.WriteMessage(messageType, []byte(err.Error())); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	switch msg.Action {
	case PUBLISH:
		if err := conn.WriteMessage(messageType, []byte(msg.Action)); err != nil {
			fmt.Println(err)
			return
		}
		// case SUBSCRIBE:
		// case UNSUBSCRIBE:

	}
}

func ChatController() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
		if err != nil {
			fmt.Println(err)
		}

		client := Client{
			Id:         autoId(),
			Connection: conn,
		}
		Clients = append(Clients, client)
		fmt.Println(len(Clients))
		payload := []byte("UUID:" + client.Id)

		conn.WriteMessage(1, payload)

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			ReceiveMessageHandler(*conn, messageType, p)
		}
	}
}
