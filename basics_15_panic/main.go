package main

import "fmt"

func process(n int){
	defer fmt.Printf("The value sent was %d\n", n)
	if n < 5 {
		panic("I panicked due to small value")
	}
	fmt.Printf("Processing value: %d\n", n)
}


func main(){
	process(10)

	process(3)
}