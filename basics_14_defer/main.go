package main

import "fmt"

// Defer is often used to ensure that resources are released, such as closing files or network connections	
// Unlocking a mutex after locking it & in debugging to print function exit messages
	
func deferprocess() {
	fmt.Println("Normal processing...")
	//defered processing
	// This statement will be executed at the end of the function
	// Defer statements are executed in LIFO order
	defer fmt.Println("first Deferred processing...")
	defer fmt.Println("second Deferred processing...")
	// This defer statement will be executed after normal processing is done
	defer fmt.Println("third Deferred processing...")
}


func defervariable() {
	// Be very careful while using variables with defer
	// The value of the variable is evaluated at the time of defer statement execution
	// not at the time of deferred function execution
	i := 10
	defer fmt.Println("Deferred value of i:", i)
	i += 5
	fmt.Println("Value of i before defer:", i)
}

func main() {
	deferprocess()
	fmt.Println("-----")
	defervariable()
}