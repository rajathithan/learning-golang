package main

import (
	"fmt"
)

func main() {
	s := "Hello, World!"
	for i, r := range s {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	// escape sequences are interpreted in double quotes
	string1 := "Hello \n World!"

	// escape sequences are not interpreted in backticks
	string2 := `Hello \n World!`

	fmt.Println("Using double quotes:{}", string1)
	fmt.Println("Using backticks:{}", string2)

}