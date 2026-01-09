package main

import (
	"errors"
	"fmt"
)

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

// functions returning multiple return of same types
func swap(x, y string) (string, string) {
	return y, x
}

// functions returning multiple return of different types
func combine(a int, b string) (string, int) {
	return b, a
}

// functions returning a type & a error type
func compare(a, b float64) (float64, error) {
	if a > b {
		return a, nil
	} else if b > a {
		return b, nil
	} else {
		//return 0, fmt.Errorf("both values are equal")
		return 0, errors.New("both values are equal")
	}
}

func main() {
	// calling the add function
	ans := add(3, 5)
	fmt.Println("The sum is:", ans)

	// passing add function as argument to operate function
	result := operate(10, 20, add)
	fmt.Println("The result of operation is:", result)

	// getting function from another function
	op := getOperation()
	// Here the op variable holds the add function
	finalResult := op(7, 8)
	fmt.Println("The final result is:", finalResult)

	// multiple return values of same types
	a, b := "hello", "world"
	x, y := swap(a, b)
	fmt.Println("Swapped values:", x, y)

	// multiple return values of different types
	str, num := combine(42, "The answer is")
	fmt.Println(str, num)

	// function returning a type & an error type
	greaterValue, err := compare(10.232323, 2.23434323232)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The greater value is :", greaterValue)
	}

}
