package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// Initialize a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Define the HTTP handler function
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Increment the wait group counter
		wg.Add(1)

		// Start a goroutine to handle the request asynchronously
		go func() {
			defer wg.Done()

			// Simulate some processing time
			processTime := 3 // in seconds
			fmt.Printf("Processing request for %d seconds...\n", processTime)
			// Simulate a time-consuming task
			for i := 0; i < processTime; i++ {
				fmt.Printf("Processing... %d seconds\n", i+1)
				// Simulate some work
				<-time.After(time.Second)
			}

			// Respond to the client after the processing is done
			fmt.Fprintln(w, "Hello, Goroutine!")

		}()
	})

	// Start the HTTP server
	go func() {
		fmt.Println("Server listening on :8080")
		http.ListenAndServe(":8080", nil)
	}()

	// Wait for all goroutines to finish before exiting the program
	wg.Wait()
}
