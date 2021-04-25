package websocket

import (
	"fmt"
)

var GlobalHash string = ""
var GlobalFlag string = ""

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan HashStruct
	ClientBroadcast chan Message
	SendHash chan string
}

func NewPool(flag string) *Pool {
	GlobalFlag = flag;
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan HashStruct),
		ClientBroadcast:  make(chan Message),
		SendHash: make(chan string),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			//fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			//for client, _ := range pool.Clients {
			//	fmt.Println(client)
			//	client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			//}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			//fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			//for client, _ := range pool.Clients {
			//	client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			//}
			break
		case HashStruct := <-pool.Broadcast:
			//fmt.Println("Broadcasting Hash to all clients!")
			//fmt.Println("HASH STRING: " + HashStruct.Hash)
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(HashStruct.Hash); err != nil {
					fmt.Println(err)
					return
				}
			}
		case message := <-pool.ClientBroadcast:
			// Send return messages to an individual client

			//fmt.Println("Sending message to Single Client")
			//fmt.Println("Client ID from Message: " + message.ClientInfo.ID)
			//if err := message.ClientInfo.Conn.WriteJSON("Thanks for the message! " + message.ClientInfo.ID); err != nil {
			//fmt.Println("message.Body = " + message.Body);
			//fmt.Println("globalhash = " + GlobalHash);
			if (message.Body == GlobalHash) {
				if err := message.ClientInfo.Conn.WriteJSON(GlobalFlag); err != nil {
					fmt.Println(err)
					return
				}
			} else {
				if err := message.ClientInfo.Conn.WriteJSON("Unknown Response"); err != nil {
					fmt.Println(err)
					return
				}
			}
			
				
		case hash := <-pool.SendHash:
			//fmt.Println("Hash Received")
			//fmt.Println("Current Hash: " + hash)
			GlobalHash = hash
		}
	}
}
