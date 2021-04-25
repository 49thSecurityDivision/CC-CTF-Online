package websocket

import (
    "fmt"
    "log"
    "github.com/gorilla/websocket"
)

type Client struct {
    ID   string
    Conn *websocket.Conn
    Pool *Pool
}

type Message struct {
    ClientInfo *Client
    Type int    `json:"type"`
    Body string `json:"body"`
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
	message := Message{ClientInfo: c, Type: messageType, Body: string(p)}
    
    fmt.Printf("Message Received: %+v\n", message.Body)
    c.Pool.ClientBroadcast <- message
    //fmt.Println("Message recieved from client: " + c.ID)
    
	}
}
