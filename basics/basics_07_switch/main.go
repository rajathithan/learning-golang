package main

import "fmt"

func main(){

	// Switch case with Multiple conditions
	day:= "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It is a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It is a weekend")
	default:
		fmt.Println("Unknown data")
	}


	// Fallthrough example
	num:= 3
	switch {
	case num > 1:
		fmt.Println("Greater than 1 but Less than 5")
		fallthrough
	case num >= 2:
		fmt.Println("Greater than 2 but Less than 5")
	default:
		fmt.Println("Odd number")
	}

	// Type Switch
	var val interface{}
	val = 42

	switch v := val.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
		// fallthrough is not allowed in type switches
	case float64:
		fmt.Printf("Float: %f\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Println("Unknown type")
	}

}