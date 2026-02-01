package main

import (
	"fmt"
)	

// generic struct definition
type Stack[T any] struct {
	elements []T
}

// Push method to add an element to the stack
func (s * Stack[T]) push (element T) {
	s.elements = append(s.elements, element)
}

// Pop method to remove and return the top element from the stack
func (s * Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false // return zero value and false if stack is empty
	}
	// Pop the last element from the slice
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element, true
}

// isEmpty method to check if the stack is empty
func (s * Stack[T]) isEmpty () bool {
	return len(s.elements) == 0
}

// printStack method to print the elements of the stack
func (s * Stack[T]) printStack () {
	fmt.Println("Stack elements:", s.elements)
}

func main() {
	// Create a stack for integers
	intStack := Stack[int]{}
	intStack.push(10)
	intStack.push(20)
	intStack.push(30)
	intStack.printStack()
	value, ok := intStack.pop()
	fmt.Println("Popped from intStack:", value, ok)
	intStack.printStack()
	value, ok = intStack.pop()
	fmt.Println("Popped from intStack:", value, ok)
	value, ok = intStack.pop()
	fmt.Println("Popped from intStack:", value, ok)
	value, ok = intStack.pop()
	fmt.Println("Popped from intStack:", value, ok)
	fmt.Println("Is intStack empty?", intStack.isEmpty())

	// Create a stack for strings
	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	fmt.Println()
	stringStack.printStack()
	strValue, _ := stringStack.pop()
	fmt.Println("Popped from stringStack:", strValue)
	stringStack.printStack()
}

