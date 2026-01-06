package main

import "fmt"

func main() {
	// Array declaration and initialization
	var fruits []string = []string{"Apple", "Banana", "Cherry"}
	fmt.Println("fruits:", fruits)

	// Slicing an array
	var numbers = [5]int{10, 20, 30, 40, 50}
	fmt.Println("numbers array:", numbers)
	var slice1 []int = numbers[1:4] // Slicing from index 1 to 3
	fmt.Println("slice1:", slice1)

	// Modifying a slice
	slice1[0] = 25
	fmt.Println("Modified slice1:", slice1)
	fmt.Println("Original array after modifying slice1:", numbers)

	
	// copying slices
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println("slice2 (copied from slice1):", slice2)

	// Modifying slice2 will not affect slice1 & original array
	slice2[0] = 99
	fmt.Println("Modified slice2:", slice2)
	fmt.Println("slice1 after modifying slice2:", slice1)
	fmt.Println("Original array after modifying slice2:", numbers)

	
	// Length and capacity of slices
	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))

	// Appending to a slice
	slice1 = append(slice1, 60, 70)
	fmt.Println("slice1 after appending:", slice1)
	fmt.Println("Length of slice1 after appending:", len(slice1))
	// When appending, if the capacity is exceeded, a new underlying array is created
	// capacity of slice is 4, but we are adding 2 more elements,
	// the final capacity will be doubled to 8
	fmt.Println("Capacity of slice1 after appending:", cap(slice1))
	fmt.Println("Original array after appending to slice1:", numbers)

}