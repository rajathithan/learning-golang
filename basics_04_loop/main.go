package main

import (
	"fmt"
)

func main() {
	// Iteration over a range of numbers from 0 to 9
	fmt.Println("Iteration over a range of numbers from 0 to 9")
		
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println()
	}

	// Iteration over a slice of strings
	words := []string{"hello", "world", "from", "Go"}
	fmt.Println("Iteration over a slice of strings")
	for _, word := range words {		
		fmt.Println(word)
		fmt.Println()
	}

	// Iteration over a map of string to int
	counts := map[string]int{"apple": 2, "banana": 3, "cherry": 5}
	fmt.Println("Iteration over a map of string to int")
	for fruit, count := range counts {		
		fmt.Printf("%s: %d\n", fruit, count)
		fmt.Println()
	}

	// Ieration over a collection
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Iteration over a collection with index and value")
	for index, number := range numbers {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
		fmt.Println()
	}

	//Iteration
	for i:= range 10{
		fmt.Println(10-i)
	}
}