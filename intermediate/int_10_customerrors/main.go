package main

import (
	"errors"
	"fmt"
)

// custom http error struct
type httpError struct {
	statusCode int
	message    string
}

// Implementing the Error method from the error interface
func (e *httpError) Error() string {
	return fmt.Sprintf("Status Code: %d, Message: %s", e.statusCode, e.message)
}

// Function to create a new httpError
func newHTTPError(code int, msg string) error {
	return &httpError{
		statusCode: code,
		message:    msg,
	}
}

func fetchResource(url string) error {
	// Simulating an error scenario
	if url == "" {
		return newHTTPError(400, "Bad Request: URL cannot be empty")
	} 
	// Simulate successful fetch
	return nil
}

func main() {
	var urls = []string{ "http://example.com", "" ,"http://golang.org"}

	for _, url := range urls {
		err := fetchResource(url)
		if err != nil {
			var httpErr *httpError
			// Using errors.As to check and extract the custom error type
			if errors.As(err, &httpErr) {
				fmt.Printf("HTTP Error occurred - %s\n", httpErr.Error())
			} else {
				fmt.Printf("An error occurred - %s\n", err.Error())
			}
		} else {
			fmt.Printf("%s - Resource fetched successfully\n", url)
		}
	}
}
