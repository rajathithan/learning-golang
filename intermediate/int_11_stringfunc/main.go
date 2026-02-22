package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello, World!"
	fmt.Println("Original string:", str)

	// Convert to uppercase
	upperStr := strings.ToUpper(str)
	fmt.Println("Uppercase string:", upperStr)

	// Convert to lowercase
	lowerStr := strings.ToLower(str)
	fmt.Println("Lowercase string:", lowerStr)

	// Check if the string contains a substring
	substr := "World"
	contains := strings.Contains(str, substr)
	fmt.Printf("Does the original string contain '%s'? %t\n", substr, contains)

	// Replace a substring
	replacedStr := strings.Replace(str, "World", "Go", 1)
	fmt.Println("Replaced string:", replacedStr)

	// Split the string into a slice
	splitStr := strings.Split(str, ", ")
	fmt.Println("Split string:", splitStr)

	// Join a slice of strings into a single string
	joinedStr := strings.Join(splitStr, " - ")
	fmt.Println("Joined string:", joinedStr)

	// Trim whitespace from the string
	trimmedStr := strings.TrimSpace("   Hello, Go!   ")
	fmt.Println("Trimmed string:", trimmedStr)

	// Get the length of the string
	length := len(str)
	fmt.Println("Length of the original string:", length)

	// Check if the string has a prefix or suffix
	hasPrefix := strings.HasPrefix(str, "Hello")
	hasSuffix := strings.HasSuffix(str, "!")
	fmt.Printf("Does the original string start with 'Hello'? %t\n", hasPrefix)
	fmt.Printf("Does the original string end with '!'? %t\n", hasSuffix)

	// Get the index of a substring
	index := strings.Index(str, "World")
	fmt.Printf("Index of 'World' in the original string: %d\n", index)	

	// Get a substring using slicing
	substring := str[7:12]
	fmt.Println("Substring from index 7 to 12:", substring)

	// Concatenate strings
	concatenatedStr := str + " Welcome to Go programming!"
	fmt.Println("Concatenated string:", concatenatedStr)

	// Repeat a string
	repeatedStr := strings.Repeat("Go! ", 3)
	fmt.Println("Repeated string:", repeatedStr)

	// Check if the string is empty
	emptyStr := ""
	isEmpty := len(emptyStr) == 0
	fmt.Printf("Is the string empty? %t\n", isEmpty)

	// Get the ASCII value of a character
	char := 'A'
	asciiValue := int(char)
	fmt.Printf("ASCII value of '%c': %d\n", char, asciiValue)

	// Get the Unicode code point of a character
	unicodeValue := rune(char)
	fmt.Printf("Unicode code point of '%c': %U\n", char, unicodeValue)	

	// Get the byte length of a string
	byteLength := len(str)
	fmt.Printf("Byte length of the original string: %d\n", byteLength)

	// Get the number of runes (characters) in a string
	runeCount := len([]rune(str))
	fmt.Printf("Number of runes in the original string: %d\n", runeCount)

	// Check if the string is a valid email address (simple check)
	email := "abc@xyz.com"
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	fmt.Printf("Is '%s' a valid email address? %t\n", email, isValidEmail)

	// convert integer to string
	num := 123
	numStr := strconv.Itoa(num)
	fmt.Printf("Integer %d converted to string: %s\n", num, numStr)

	// Repeat a string multiple times
	repeatedStr2 := strings.Repeat("Go! ", 5)
	fmt.Println("Repeated string 2:", repeatedStr2)

	// Count character occurrences in a string
	charToCount := 'o'
	count := strings.Count(str, string(charToCount))
	fmt.Printf("Number of occurrences of '%c' in the original string: %d\n", charToCount, count)

	// Check if it has a prefix or suffix
	hasPrefix2 := strings.HasPrefix(str, "Hello")
	hasSuffix2 := strings.HasSuffix(str, "!")
	fmt.Printf("Does the original string start with 'Hello'? %t\n", hasPrefix2)
	fmt.Printf("Does the original string end with '!'? %t\n", hasSuffix2)

	// find digits in a string using regexp	
	str4 := "The price is 100 dollars and 50 cents."
	re := regexp.MustCompile("[0-9]+")
	// -1 means find all occurrences
	digits := re.FindAllString(str4, -1)
	fmt.Printf("Digits in the original string: %v\n", digits)

	// Counts chars in other language strings using utf8 rune count	
	str5 := "你好，世界！"
	runeCount2 := utf8.RuneCountInString(str5)
	fmt.Printf("Number of runes in the string '%s': %d\n", str5, runeCount2)

	// string builder for efficient string concatenation
	// string build is performance efficient when concatenating multiple strings as it minimizes memory copying
	var builder strings.Builder

	// WriteString is used to append strings to the builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")
	finalString := builder.String()
	fmt.Println("String built using strings.Builder:", finalString)

	// WriteRune is used to append a single rune (character) to the builder
	builder.WriteRune(' ')
	builder.WriteRune('G')
	builder.WriteRune('o')
	finalString2 := builder.String()
	fmt.Println("String built using strings.Builder with runes:", finalString2)

	// Reset the builder to reuse it
	builder.Reset()
	builder.WriteString("Go is great!")
	finalString3 := builder.String()
	fmt.Println("String built after resetting the builder:", finalString3)
}
