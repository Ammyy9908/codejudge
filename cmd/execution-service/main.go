package main

import (
	"log"

	"github.com/ammyy9908/codejudge/internal/execution/queue"
)

func main() {
	log.Println("Starting execution-service...")

	// Start consuming submissions
	err := queue.StartConsumer()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
}
