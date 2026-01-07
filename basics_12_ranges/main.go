package main

import "fmt"

func main() {
	var message = "Hello, Ranges!"
	for i, v := range message {
		// Prints index and Decimal rune of each character
		fmt.Println(i,v)
		
		// Prints index and Unicode of each character
		fmt.Printf("Index :{%d}, Unicode:{%U}\n", i, v)

		// Prints index and rune of each character
		fmt.Printf("Index i - %d, Rune - %c\n", i, v)
	}
}
