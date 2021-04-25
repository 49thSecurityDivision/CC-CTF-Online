package main

import (
	"fmt"
	"net/http"
	"os"
	"websocket"

	"github.com/google/uuid"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	//fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		ID:   uuid.New().String(),
		Conn: conn,
		Pool: pool,
	}

	// Add the new client to the pool
	pool.Register <- client
	client.Read()
}

func setupRoutes(word string, flag string) {

	pool := websocket.NewPool(flag)
	go pool.Start()
	go websocket.GenHash(pool, word)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
		//fmt.Fprintf(w, "Hello World")
	})

	wordPath := "/" + word
	http.HandleFunc(wordPath, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(flag))
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

}

func main() {

	var word, flag string
	if len(os.Args) != 3 {
		//fmt.Println("Usage:", os.Args[0], "Secret_Word", "Flag");
		fmt.Println("No arguments given... checking for ENV...")

		if len(os.Getenv("WORD")) == 0 || len(os.Getenv("FLAG")) == 0 {
			os.Exit(1)
		} else {
			word = os.Getenv("WORD")
			flag = os.Getenv("FLAG")
		}

	} else {
		word = os.Args[1]
		flag = os.Args[2]
	}

	//go websocket.Start()
	setupRoutes(word, flag)
	//go hashgen.GenHash()
	http.ListenAndServe(":8080", nil)
}
