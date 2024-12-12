// File: backend/internal/websocket/handlers/handlers.go
package handlers

import (
	"log"
	"net/http"

	redis "github.com/ammyy9908/codejudge/internal/websocket/redis"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin for simplicity
	},
}

// HandleWebSocket manages WebSocket connections from clients
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	submissionID := r.URL.Query().Get("submission_id")
	if submissionID == "" {
		http.Error(w, "submission_id query param is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade WebSocket connection: %v", err)
		return
	}

	// Register WebSocket connection in Redis
	if err := redis.RegisterConnection(submissionID, conn); err != nil {
		log.Printf("Failed to register WebSocket connection: %v", err)
		conn.Close()
		return
	}

	log.Printf("WebSocket connection established for submission_id: %s", submissionID)

	// Keep the connection alive
	defer redis.RemoveConnection(submissionID)
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			log.Printf("WebSocket connection closed for submission_id: %s", submissionID)
			break
		}
	}
}
