package queue

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	execution "github.com/ammyy9908/codejudge/internal/execution/models"
)

var executionQueue = "execution_queue"

// PublishExecutionResult sends the execution result to RabbitMQ
func PublishExecutionResult(result execution.ExecutionResult) error {
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

	// Declare the queue
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

	// Publish the execution result
	body, _ := json.Marshal(result)
	err = ch.Publish(
		"",
		executionQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Execution result published: %s", result.ID)
	return nil
}
