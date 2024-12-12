package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	connections map[string]*websocket.Conn // Maps submission_id to WebSocket connections
	mu          sync.RWMutex               // Protects the connections map
}

var manager = &ConnectionManager{
	connections: make(map[string]*websocket.Conn),
}

// RegisterConnection adds a WebSocket connection for a given submission_id
func RegisterConnection(submissionID string, conn *websocket.Conn) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	manager.connections[submissionID] = conn
}

// GetConnection retrieves the WebSocket connection for a given submission_id
func GetConnection(submissionID string) (*websocket.Conn, bool) {
	manager.mu.RLock()
	defer manager.mu.RUnlock()
	conn, exists := manager.connections[submissionID]
	return conn, exists
}

// RemoveConnection removes a WebSocket connection from the manager
func RemoveConnection(submissionID string) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	delete(manager.connections, submissionID)
}
