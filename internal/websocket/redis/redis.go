// File: backend/internal/websocket/redis/redis.go
package redis

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var (
	rdb  *redis.Client
	ctx  = context.Background()
	lock sync.RWMutex

	connections = make(map[string]*websocket.Conn) // Local cache for WebSocket connections
)

// InitializeRedis sets up the Redis client
func InitializeRedis(addr, password string) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // Default DB
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return err
	}

	log.Println("Redis connection established")
	return nil
}

// RegisterConnection adds a WebSocket connection to the local cache and Redis
func RegisterConnection(submissionID string, conn *websocket.Conn) error {
	lock.Lock()
	defer lock.Unlock()

	connections[submissionID] = conn
	if err := rdb.Set(ctx, submissionID, "active", 0).Err(); err != nil {
		return err
	}

	return nil
}

// GetConnection retrieves a WebSocket connection from the local cache
func GetConnection(submissionID string) (*websocket.Conn, error) {
	lock.RLock()
	defer lock.RUnlock()

	conn, exists := connections[submissionID]
	if !exists {
		return nil, errors.New("connection not found")
	}

	return conn, nil
}

// RemoveConnection removes a WebSocket connection from the local cache and Redis
func RemoveConnection(submissionID string) error {
	lock.Lock()
	defer lock.Unlock()

	if conn, exists := connections[submissionID]; exists {
		conn.Close()
		delete(connections, submissionID)
	}

	if err := rdb.Del(ctx, submissionID).Err(); err != nil {
		return err
	}

	return nil
}
