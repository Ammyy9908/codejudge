package queue

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	execution "github.com/ammyy9908/codejudge/internal/execution/models"
	"github.com/ammyy9908/codejudge/internal/execution/sandbox"
)

var rabbitMQURL = "amqp://guest:guest@localhost:5672/"
var submissionQueue = "submission_queue"

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

	// Declare the queue
	_, err = ch.QueueDeclare(
		submissionQueue,
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
		submissionQueue,
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
		var submission execution.Submission
		if err := json.Unmarshal(msg.Body, &submission); err != nil {
			log.Printf("Failed to unmarshal submission: %v", err)
			continue
		}

		// Execute the code
		result := sandbox.ExecuteCode(submission)

		// Publish the result
		if err := PublishExecutionResult(result); err != nil {
			log.Printf("Failed to publish execution result: %v", err)
		}
	}

	return nil
}
