package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {

	fmt.Println("websocket server started")

	//create a Room that client can connect
	//client recieves a message from client and send it to all clients in the Room

	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8080", nil)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("messageType:", messageType)
		fmt.Println("p:", p)
		fmt.Println("String:", string(p[:]))
	}

}
