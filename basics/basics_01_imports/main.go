package main

import (
	"fmt"
	// Importing the net/http package with an alias 'foo'
	foo "net/http"
)

func main() {
	fmt.Println("Import example")
	// Using the aliased package 'foo' to make a GET request
	response, err := foo.Get("http://example.com")
	// Handling error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Ensuring the response body is closed after we're done with it
	defer response.Body.Close()
	// Printing the response status
	fmt.Println("Response Status:", response.Status)
}