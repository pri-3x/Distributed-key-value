package main

import (
	"sync"
)

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewStore initializes a new key-value store
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Put stores a key-value pair
func (s *Store) Put(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get retrieves a value by key
func (s *Store) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}
