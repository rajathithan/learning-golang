package main

import "fmt"

var firstname string = "Rajathithan"

func main() {
	// use of variable before declaration
	fmt.Println("Printing the global variable's value :",firstname)

	// variable declaration and initialization
	firstname := "Raj"

	// use of variable after declaration
	fmt.Println("Printing the local variable's value :",firstname)
}