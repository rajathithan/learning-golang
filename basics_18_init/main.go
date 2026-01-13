package main

import "fmt"

// Init functions are always executed before the main function
// This is to configure settings, initialize variables, or perform setup tasks
// Multiple init functions will be executed in sequential order
func init() {
	fmt.Println("This is the init function1, executed before main.")
}

func init() {
	fmt.Println("This is the init function2, executed before main.")
}

func init() {
	fmt.Println("This is the init function3, executed before main.")
}

// This main function will be executed after all init functions
func main() {
	fmt.Println("This is the main function.")
}