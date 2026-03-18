package main

import (
	"fmt"
)

func main() {
	name := "Alice"
	age := 30
	height := 1.75

	// Using fmt.Sprintf to format a string
	formattedStr := fmt.Sprintf("Name: %s, Age: %d, Height: %.2f meters", name, age, height)
	fmt.Println(formattedStr)

	// Using fmt.Printf to print formatted output directly
	fmt.Printf("Name: %s, Age: %d, Height: %.2f meters\n", name, age, height)

	// Formatting numbers with leading zeros
	num1 := 345
	num2 := 12345
	num3 := 123456

	fmt.Printf("%05d\n",num1)
	fmt.Printf("%05d\n",num2)
	fmt.Printf("%05d\n",num3)

	// Formatting strings with width and alignment
	message := "Hello"
	fmt.Printf("|%10s|\n",message)
	fmt.Printf("|%-10s|\n",message)

	// Using raw string literals to preserve formatting
	str1 := "Hello \nWorld"
	fmt.Println(str1)

	// Using raw string literals to preserve formatting
	str2 := `Hello \nWorld`
	fmt.Println(str2)
}
