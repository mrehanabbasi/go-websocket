package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
		if err := conn.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(5*time.Second)); err != nil {
			log.Println(err)
			return
		} else {
			log.Println("Websocket Connection Closed!")
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// var rh = http.Header{
	// 	"Content-Type": {"application/json"},
	// }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Client successfully connected!!!")

	reader(ws)
	// defer ws.Close()
}

func setupRoutes() {
	http.HandleFunc("/", wsEndpoint)
}

func main() {
	fmt.Println("Webhooks!!!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
