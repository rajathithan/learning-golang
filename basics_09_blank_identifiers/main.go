package main

import "fmt"

func getValues() (int, int) {
	return 10, 20
}

func main() {
	// Blank identifier example
	value1, _ := getValues()
	_, value2 := getValues()
	fmt.Println("Value1:", value1)
	fmt.Println("Value2:", value2)

	// Ignoring multiple return values
	_, _ = getValues()	

	// Using blank identifier in loops
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum of numbers:", sum)
}

