package main

import (
	"fmt"
	"log"
	"net/http"


	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("📩 Received:", string(message))
		// echo back
		conn.WriteMessage(websocket.TextMessage, []byte("Pong from server"))
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket server started on :5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
