package main

import "fmt"

func main() {
	score := 75
	
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	a:= 11
	b:= 21

	if a % 2 == 0  && b % 2 == 0 {
		fmt.Println("Both a and b are even")
	} else if a % 2 == 0 || b % 2 == 0 {
		fmt.Println("Either a or b is even")
	} else {
		fmt.Println("Both a and b are odd")
	}
}
