package main

import "fmt"


func process(){
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	
	fmt.Println("Starting process...")
	panic("Something went wrong!")
	// Any code after panic will not be executed
}


func main(){
	process()
	
}