package main

import (
	"sync"
)

type Node struct {
	ID     string
	Addr   string
	Status bool // true for healthy, false for failed
}

type Cluster struct {
	nodes map[string]*Node
	mu    sync.RWMutex
}

// NewCluster initializes a new cluster
func NewCluster() *Cluster {
	return &Cluster{
		nodes: make(map[string]*Node),
	}
}

// AddNode adds a new node to the cluster
func (c *Cluster) AddNode(node Node) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nodes[node.ID] = &node
}

// RemoveNode removes a node from the cluster
func (c *Cluster) RemoveNode(nodeID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.nodes, nodeID)
}

// HealthCheck checks the status of nodes and updates their status
func (c *Cluster) HealthCheck() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, node := range c.nodes {
		// Here you would implement actual health check logic
		// For demonstration, we will just print the node status
		if node.Status {
			// Simulate a health check
			node.Status = true
		} else {
			node.Status = false
		}
	}
}

// GetNodes returns a list of all nodes in the cluster
func (c *Cluster) GetNodes() []*Node {
	c.mu.RLock()
	defer c.mu.RUnlock()
	nodes := make([]*Node, 0, len(c.nodes))
	for _, node := range c.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}
