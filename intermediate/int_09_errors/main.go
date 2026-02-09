package main

import (
	"errors"
	"fmt"
	"math"
)

type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

func newCustomError(msg string) error {
	return &customError{message: msg}
}

func sqrt(value float64) (float64, error) {
	if value < 0 {
		// Implementing error using the standard errors package
		return 0, errors.New("cannot compute square root of a negative number")

		// Implementing error using a custom error type
		//return 0, newCustomError("cannot compute square root of a negative number")

		// Implementing error using fmt.Errorf for formatted messages
		//return 0, fmt.Errorf("cannot compute square root of %v", value)
	}
	return math.Sqrt(value), nil
}

func process_data(data []float64) {
	for _, v := range data {
		if result, err := sqrt(v); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Square root of %v is %v\n", v, result)
		}
	}
}

func main() {
	values := []float64{16, 25, -4, 9}
	process_data(values)

}
