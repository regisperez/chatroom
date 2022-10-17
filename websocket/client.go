package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		timeMessage:=time.Now().Format("2006-01-02 15:04:05")
		message := Message{Type: messageType, Body: timeMessage+" "+c.ID+": " +string(p)}
		c.Pool.Broadcast <- message
	}
}
