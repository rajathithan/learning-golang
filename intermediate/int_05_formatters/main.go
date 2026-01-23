package main

import (
	"fmt"
)

func main() {
	// Using different format specifiers in Printf
	// %v - default format
	// %#v - Go-syntax representation
	// %T - type of the variable
	// %% - literal percent sign

	floatNum := 2_55.2343356

	fmt.Printf("Default format: %v\n", floatNum)
	fmt.Printf("Go-syntax representation: %#v\n", floatNum)
	fmt.Printf("Type of variable: %T\n", floatNum)
	fmt.Printf("Literal percent sign: %v %%\n", floatNum)
}
