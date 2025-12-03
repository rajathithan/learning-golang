package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Import example")
	response, err := foo.Get("http://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	
	fmt.Println("Response Status:", response.Status)
}