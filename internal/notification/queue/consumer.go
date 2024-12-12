// File: internal/notification/queue/consumer.go
package queue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	execution "github.com/ammyy9908/codejudge/internal/execution/models"
	"github.com/streadway/amqp"
)

var rabbitMQURL = getRabbitMQURL()

func getRabbitMQURL() string {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		return "amqp://guest:guest@localhost:5672/"
	}
	return url
}

var websocketServiceURL = getWebSocketServiceURL()

func getWebSocketServiceURL() string {
	url := os.Getenv("WEBSOCKET_SERVICE_URL")
	if url == "" {
		return "http://localhost:8000/send" // Default for local development
	}
	return url
}

var (
	executionQueue = "execution_queue"
)

// StartConsumer starts consuming execution results from the RabbitMQ execution queue
func StartConsumer() error {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open RabbitMQ channel: %v", err)
		return err
	}
	defer ch.Close()

	// Declare the execution queue
	_, err = ch.QueueDeclare(
		executionQueue,
		true,  // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Printf("Failed to declare queue: %v", err)
		return err
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		executionQueue,
		"",
		true,  // Auto-acknowledge
		false, // Non-exclusive
		false, // No-local
		false, // No-wait
		nil,
	)
	if err != nil {
		log.Printf("Failed to start consuming: %v", err)
		return err
	}

	for msg := range msgs {
		var result execution.ExecutionResult
		if err := json.Unmarshal(msg.Body, &result); err != nil {
			log.Printf("Failed to unmarshal execution result: %v", err)
			continue
		}

		// Forward the result to the WebSocket service
		if err := forwardToWebSocketService(result); err != nil {
			log.Printf("Failed to forward result to WebSocket service: %v", err)
		}
	}

	return nil
}

// forwardToWebSocketService forwards execution results to the WebSocket service
func forwardToWebSocketService(result execution.ExecutionResult) error {
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	resp, err := http.Post(websocketServiceURL, "application/json", bytes.NewBuffer(resultJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("WebSocket service returned status %d", resp.StatusCode)
	}

	log.Printf("Result forwarded to WebSocket service for submission_id: %s", result.ID)
	return nil
}
