package main

import "fmt"

func main() {
	fmt.Println("Fibonacci Sequence:", fibonacci(3))
	fmt.Println("Fibonacci Sequence:", fibonacci(15))
}

// fibonacci returns the fibonacci sequence using recursion
func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}

	seq := fibonacci(n - 1)
	nextValue := seq[len(seq)-1] + seq[len(seq)-2]
	return append(seq, nextValue)
}
	