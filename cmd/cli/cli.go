package main

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
)

func main() {
	url := "ws://localhost:8080/ws"
	log.Printf("Connecting to %s...", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Ожидание выхода
	<-interrupt
	log.Println("Interrupted. Closing connection.")
	conn.Close()
}
