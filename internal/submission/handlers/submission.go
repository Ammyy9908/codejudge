package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ammyy9908/codejudge/internal/submission/models"
	"github.com/ammyy9908/codejudge/internal/submission/queue"
)

// SubmitHandler handles HTTP POST requests for code submissions
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON payload from the request body
	var submission models.Submission
	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// // Generate a unique submission ID
	// submission.ID = utils.GenerateUUID()

	// Enqueue the task in RabbitMQ
	if err := queue.PublishSubmission(submission); err != nil {
		log.Printf("Failed to enqueue submission: %v", err)
		http.Error(w, "Failed to process submission", http.StatusInternalServerError)
		return
	}

	// Send JSON response back to the client
	w.WriteHeader(http.StatusAccepted)
	response := map[string]string{
		"submission_id": submission.ID,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
