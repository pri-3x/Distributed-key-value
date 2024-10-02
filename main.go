package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Initialize the key-value store
	store := NewStore()

	// Start the HTTP server for the client interface
	http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		store.Put(key, value)
		fmt.Fprintf(w, "Stored %s: %s\n", key, value)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := store.Get(key)
		fmt.Fprintf(w, "Value for %s: %s\n", key, value)
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
