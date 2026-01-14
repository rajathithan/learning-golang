package main

import "fmt"

func main(){
	//closure is a special kind of function value (often an anonymous function) that 
	// "closes over" or remembers the variables from its surrounding lexical scope, 
	// even after the outer function has finished executing
	// This allows functions to maintain state between calls 
	// without resorting to global variables or complex object-oriented structures. 

	adder := add()
	fmt.Println(adder(3)) // Output: 3
	fmt.Println(adder(5)) // Output: 8
	fmt.Println(adder(10)) // Output: 18

	// Each call to add() creates a new instance of the closure with its own state
	anotherAdder := add()
	fmt.Println(anotherAdder(2)) // Output: 2
	fmt.Println(anotherAdder(4)) // Output: 6

	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // Output: 10
	fmt.Println(triple(5)) // Output: 15
}

func add() func(int) int {
	fmt.Println("Initializaing the sum to 0 !!")
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}