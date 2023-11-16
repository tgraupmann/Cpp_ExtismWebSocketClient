package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	serverAddr := "ws://localhost:8080/ws" // Change this to your WebSocket server address

	// Establish WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to WebSocket server")

	// Start a goroutine to read and print messages from the server
	go readMessages(conn)

	// Wait for 10 seconds and then exit
	timer := time.NewTimer(10 * time.Second)
	<-timer.C
	fmt.Println("Exiting after 10 seconds.")
}

func readMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		fmt.Println("Received message:", string(message))
	}
}
