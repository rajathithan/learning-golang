package main

import "fmt"

func doubleValue(p *int) {
	*p = *p * 2
}

func main() {

	// Pointer declaration and usage
	var ptr *int
	var num int = 10
	// Assigning the address of num to ptr
	ptr = &num

	fmt.Println("The memory address of num is:", ptr)
	// deferencing the pointer to get the value of num
	fmt.Println("The current value of num is:", *ptr)
	// Modifying the value of num using the pointer
	doubleValue(ptr)
	// The address remains the same after modification
	fmt.Println("The memory address of num after doubling value is:", ptr)
	fmt.Println("The current value of num is:", *ptr)
}

