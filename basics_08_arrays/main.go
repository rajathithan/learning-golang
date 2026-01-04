package main

import "fmt"

func main() {
	// Array declaration and initialization
	// var arrayName [size]dataType
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", numbers)

	// Accessing array elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Third element:", numbers[2])

	// Modifying array elements
	numbers[1] = 25
	fmt.Println("Modified Array:", numbers)

	// Length of the array
	fmt.Println("Length of the array:", len(numbers))

	// Iterating over an array
	fmt.Println("Array elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Using range to iterate over an array
	fmt.Println("Array elements using range:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}	

	// Multi-dimensional array
	var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multi-dimensional Array:", matrix)

	// Accessing elements in multi-dimensional array
	fmt.Println("Element at row-2 & Column-3 - (1,2):", matrix[1][2])

	// Iterating over multi-dimensional array
	fmt.Println("Multi-dimensional Array elements:")
	// Iterate each row
	for i := 0; i < len(matrix); i++ {
		// Iterate each column
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at (%d,%d): %d\n", i, j, matrix[i][j])
		}
	}

	//copying arrays
	var src = [3]int{1, 2, 3}
	dest := src // copies the entire array
	dest[0] = 10
	fmt.Println("Source Array is not modified:", src)
	fmt.Println("Only destination array is modified:", dest)

	// Pointer to array
	var src2 = [3]int{7, 8, 9}
	// dest2 is a pointer to an array of 3 integers
	var dest2 *[3]int = &src2
	dest2[0] = 10
	fmt.Println("Source Array is modified due to pointer modification:", src2)
	

	// Array comparison
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}

	fmt.Println("arr1 == arr2:", arr1 == arr2) // true
	fmt.Println("arr1 == arr3:", arr1 == arr3) // false

}
