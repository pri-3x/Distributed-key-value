package main

import (
	"fmt"
	"sync"
)

type ConflictResolver struct {
	mu sync.Mutex
}

// HandleConflict resolves conflicts between nodes
func (cr *ConflictResolver) HandleConflict(key string, values []string) string {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	// Simple conflict resolution: return the first non-empty value
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

// RedirectRequest handles requests to failed nodes
func RedirectRequest(nodeID string) {
	// In a real implementation, you would implement logic to redirect requests
	// For demonstration, we will just print a message
	fmt.Printf("Redirecting request from failed node: %s\n", nodeID)
}

// LogError logs errors for debugging purposes
func LogError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
