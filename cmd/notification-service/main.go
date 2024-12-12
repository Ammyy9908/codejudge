package main

import (
	"log"

	"github.com/ammyy9908/codejudge/internal/notification/queue"
)

func main() {
	log.Println("Starting notification-service...")

	// Start consuming execution results
	err := queue.StartConsumer()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
}
