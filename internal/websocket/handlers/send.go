package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	execution "github.com/ammyy9908/codejudge/internal/execution/models"
	redis "github.com/ammyy9908/codejudge/internal/websocket/redis"
)

// SendResult handles incoming results and forwards them to WebSocket clients
func SendResult(w http.ResponseWriter, r *http.Request) {
	var result execution.ExecutionResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Get the WebSocket connection for the submission_id
	conn, exists := redis.GetConnection(result.ID)
	if exists != nil {
		log.Printf("No active WebSocket connection found for submission_id: %s", result.ID)
		http.Error(w, "WebSocket connection not found", http.StatusNotFound)
		return
	}

	// Send the result to the WebSocket client
	if err := conn.WriteJSON(result); err != nil {
		log.Printf("Failed to send result to WebSocket client: %v", err)
		http.Error(w, "Failed to send result", http.StatusInternalServerError)
		return
	}

	log.Printf("Result sent to WebSocket client for submission_id: %s", result.ID)
	w.WriteHeader(http.StatusOK)
}
