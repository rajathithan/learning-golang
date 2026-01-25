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

	// string formatting verbs
	str := "Gopher"

	fmt.Printf("String: %s\n", str)
	fmt.Printf("String with quotes: %q\n", str)
	fmt.Printf("String with width 10 right aligned: %10s\n", str)
	fmt.Printf("String with width 10 left aligned: %-10s\n", str)
	fmt.Printf("String with width 10 and cut 3 chars: %10.3s\n", str)
	fmt.Printf("String with 3 chars: %.3s\n", str)
	fmt.Printf("String with Hexadecimal: %x\n", str)
	fmt.Printf("String with Hexadecimal space: % x\n", str)

	// float formatting verbs
	floatVal := 123.456789
	floatNum2 := 123456789.123

	fmt.Printf("Float default: %f\n", floatVal)
	fmt.Printf("Float with 2 decimal places: %.2f\n", floatVal)
	fmt.Printf("Float with width 10 and 2 decimal places: %10.2f\n", floatVal)
	fmt.Printf("Float with width 10, left aligned and 2 decimal places: %-10.2f\n", floatVal)
	fmt.Printf("Float uses exponent when needed : %g\n",floatNum2)
	fmt.Printf("Float in scientific notation: %e\n", floatVal)
	fmt.Printf("Float in scientific notation with uppercase E: %E\n", floatVal)

	// boolean formatting verbs
	boolVal := true
	
	fmt.Printf("Boolean default: %t\n", boolVal)
	fmt.Printf("Boolean with width 5 right aligned: %5t\n", boolVal)
	fmt.Printf("Boolean with width 5 left aligned: %-5t\n", boolVal)
}
