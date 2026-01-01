package main

import "fmt"

func main() {
	// for init; condition; post {}
	i:= 1;
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	for{
		fmt.Println("infinite loop")
		break
	}

	// for init; condition; post {}
	i = 1
	for{
		if i > 15 {
			break
		}
		fmt.Println(i)
		i++
	}
}