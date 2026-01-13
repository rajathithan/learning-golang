package main

import (
	"fmt"
	"os"
)

func main (){
	defer fmt.Println("Even though I am deferred, I will not be executed!")
	fmt.Println("Exiting the program now.")
	// Use exit to terminate the program immediately
	// Use this option sparingly, in cases you want to terminate the program without running deferred calls
	os.Exit(1)
	fmt.Println("This line will also never be printed.")
}