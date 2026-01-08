package main

import "fmt"

// functions are first-class citizens in Go
// you can pass them as arguments and return them from other functions

// function with two integer parameters and an integer return type
func add(a int, b int) int {
	return a + b
}

// passing function as argument
func operate(a int, b int, operation func(int, int) int) int {
	// Notice that we are calling the add function as operation here
	return operation(a, b)
}

// returning function from another function
func getOperation() func(int, int) int {
	// Returning the add function itself
	return add
}

func main() {
	ans := add(3, 5)
	fmt.Println("The sum is:", ans)

	result := operate(10, 20, add)
	fmt.Println("The result of operation is:", result)

	op := getOperation()
	// Here the op variable holds the add function
	finalResult := op(7, 8)
	fmt.Println("The final result is:", finalResult)
}