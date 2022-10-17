package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"strings"
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
		var body string
		if strings.Contains(string(p),"Stock Bot:"){
			body = timeMessage+" "+ string(p)
		}else{
			body = timeMessage+" "+c.ID+": " +string(p)
		}
		message := Message{Type: messageType, Body: body}
		c.Pool.Broadcast <- message
	}
}
