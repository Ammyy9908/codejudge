package websocket

import (
	"encoding/json"
	"log"

	execution "github.com/ammyy9908/codejudge/internal/execution/models"
	"github.com/gorilla/websocket"
)

// NotifyClient sends the execution result to the WebSocket client
func NotifyClient(result execution.ExecutionResult) error {
	log.Printf("Notifying client for submission_id: %s", result.ID)
	conn, exists := GetConnection(result.ID)
	if !exists {
		log.Printf("No active WebSocket connection found for submission_id: %s", result.ID)
		return nil
	}

	defer RemoveConnection(result.ID) // Clean up connection after sending
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, resultJSON)
	if err != nil {
		log.Printf("Failed to send message to WebSocket: %v", err)
		return err
	}

	log.Printf("Sent result to WebSocket: %s", result.ID)
	return nil
}
