package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func joinChat(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	fmt.Println("Got Inside Function")

	for {
		msgType, msg, err := conn.ReadMessage() 
		if err != nil {
			fmt.Println("Connection failed: ", err.Error())
			return
		}

		fmt.Printf("%d %s", msgType, msg)

		if err = conn.WriteMessage(msgType, msg); err != nil {
			fmt.Println("Send Failed: ", err.Error())
			return 
		}
	}
}

func homeChat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}


func main() {
	http.HandleFunc("/join", joinChat)
	http.HandleFunc("/", homeChat)

	http.ListenAndServe(":5000", nil)
}
