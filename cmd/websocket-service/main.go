package main

import (
	"log"
	"net/http"

	handlers "github.com/ammyy9908/codejudge/internal/websocket/handlers"
	redis "github.com/ammyy9908/codejudge/internal/websocket/redis"
)

func main() {
	// Initialize Redis
	if err := redis.InitializeRedis("redis:6379", ""); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Register WebSocket handler for client connections
	http.HandleFunc("/ws", handlers.HandleWebSocket)

	// Register the /send endpoint for receiving execution results
	http.HandleFunc("/send", handlers.SendResult)

	log.Println("WebSocket service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
