package queue

import (
	"encoding/json"
	"log"

	"github.com/ammyy9908/codejudge/internal/submission/models"
	"github.com/streadway/amqp"
)

// PublishSubmission publishes a submission to RabbitMQ
func PublishSubmission(submission models.Submission) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	queueName := "submission_queue"
	_, err = ch.QueueDeclare(
		queueName,
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

	body, _ := json.Marshal(submission)
	err = ch.Publish(
		"",
		queueName,
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

	log.Printf("Submission enqueued with ID: %s", submission.ID)
	return nil
}
