package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("cannot compute square root of a negative number")
	}
	return math.Sqrt(value), nil
}

func main() {
	values := []float64{16, 25, -4, 9}

	for _, v := range values {
		result, err := sqrt(v)
		if err != nil {
			fmt.Printf("Error computing sqrt(%v): %v\n", v, err)
		}	 else {
			fmt.Printf("The square root of %v is %v\n", v, result)
		}
	}
}
