package goWithWebSockets

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/extism/go-pdk"
)

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

//export _start
func _start() int32 {
	input := pdk.Input()
	greeting := `Hello, ` + string(input) + `!`
	pdk.OutputString(greeting)

	// WebSocket server address
	serverAddr := "ws://localhost:8080/ws"

	// Establish WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to WebSocket server")

	// Start a goroutine to read and print messages from the server
	go readMessages(conn)

	// Wait for interruption (Ctrl+C) to gracefully close the connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case <-interrupt:
		fmt.Println("Interrupt received, closing connection.")
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		return 0
	}

	return 0
}

func main() {}
