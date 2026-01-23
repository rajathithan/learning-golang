package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, World!"
	for i, r := range s {
		fmt.Printf("Index: %d, ASCII: %v, Rune: %c\n", i, r, r)
	}

	// escape sequences are interpreted in double quotes
	string1 := "Hello \n World!"

	// escape sequences are not interpreted in backticks
	string2 := `Hello \n World!`

	fmt.Println("Using double quotes:{}", string1)
	fmt.Println("Using backticks:{}", string2)

	// comparing strings
	str1 := "Hello" // ASCII value of 'H' is 72
	str2 := "hello" // ASCII value of 'h' is 104

	fmt.Println(str1 > str2)

	// count characters in a string
	str3 := "Hello, 世界"
	fmt.Println("Number of characters:", len([]rune(str3)))
	fmt.Println("Number of characters:", utf8.RuneCountInString(str3))

	// concatenate strings
	str4 := "Hello, "
	str5 := "Gophers!"
	concatenated := str4 + str5
	fmt.Println(concatenated)

	// unicode characters using runes
	var char1 rune = '💼'
	var char2 rune = '🥱'
	fmt.Printf("Character: %c, ASCII: %d\n", char1, char1)
	fmt.Printf("Character: %c, ASCII: %d\n", char2, char2)


}