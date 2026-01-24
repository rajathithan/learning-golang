package main

import (
	"fmt"
)

func main() {
	// Using different format specifiers in Printf
	// %v - default format
	// %#v - Go-syntax representation
	// %T - type of the variable
	// %% - literal percent sign

	floatNum := 2_55.2343356

	fmt.Printf("Default format: %v\n", floatNum)
	fmt.Printf("Go-syntax representation: %#v\n", floatNum)
	fmt.Printf("Type of variable: %T\n", floatNum)
	fmt.Printf("Literal percent sign: %v %%\n", floatNum)

	// Formatting integers in different bases
	intNum := 255
	
	fmt.Printf("Decimal: %d\n", intNum)
	fmt.Printf("Decimal with sign: %+d\n", intNum)
	fmt.Printf("Decimal right  : %4d\n", intNum)
	fmt.Printf("Decimal left   : %-4d\n", intNum)
	fmt.Printf("Decimal with zeroes padded : %04d\n", intNum)
	fmt.Printf("Binary: %b\n", intNum)
	fmt.Printf("Hexadecimal: %x\n", intNum)
	fmt.Printf("Hexadecimal with uppercase: %X\n", intNum)
	fmt.Printf("Hexadecimal with prefix: %#x\n", intNum)
	fmt.Printf("Octal: %o\n", intNum)
	fmt.Printf("Octal with prefix: %#o\n", intNum)
}
