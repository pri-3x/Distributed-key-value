package main

import (
	"fmt"
	"net/http"
)

type Client struct {
	baseURL string
}

// NewClient initializes a new client
func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

// Put sends a put request to the server
func (c *Client) Put(key, value string) {
	resp, err := http.Get(fmt.Sprintf("%s/put?key=%s&value=%s", c.baseURL, key, value))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
}

// Get sends a get request to the server
func (c *Client) Get(key string) string {
	resp, err := http.Get(fmt.Sprintf("%s/get?key=%s", c.baseURL, key))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()
	var value string
	fmt.Fscanf(resp.Body, "Value for %s: %s\n", &key, &value)
	return value
}
