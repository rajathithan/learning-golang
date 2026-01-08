
# learning-golang

Each folder contains a `main.go` file with example code and comments for easy understanding.

## 1. basics_01_imports
Demonstrates how to import packages in Go, including the use of aliases. The program makes an HTTP GET request using an aliased import for the `net/http` package and prints the response status.

## 2. basics_02_variables
Shows variable declaration, initialization, and scope. It prints a global variable, then declares a local variable with the same name and prints its value, illustrating variable shadowing in Go.

## 3. basics_03_constants
Introduces the use of constants in Go. The program defines a constant `pi` and prints its value.

## 4. basics_04_loop
Explains different types of loops in Go. The program demonstrates:
- Iterating over a range of numbers
- Iterating over a slice of strings
- Iterating over a map
- Iterating over a collection with both index and value

## 5. basics_05_for_as_while
Explores Go's `for` used as a while-style and infinite loop with `break`. The program:
- Counts from 1 to 10 using a condition-controlled `for`.
- Runs an infinite loop and exits with `break`.
- Demonstrates a manual loop that increments `i` and stops when `i > 15`.

## 6. basics_06_if_else
Demonstrates conditional logic in Go using `if`, `else if`, and `else`, along with relational and logical operators. The program:
- Categorizes a numeric `score` into letter grades (`A`, `B`, `C`, `F`).
- Checks parity of two integers (`a`, `b`) and prints whether both are even, one is even, or both are odd using `&&` and `||`.

## 7. basics_07_switch
Illustrates Go's `switch` statement with various features. The program:
- Uses a switch with multiple conditions to classify days as weekdays or weekends.
- Demonstrates `fallthrough` to execute the next case.
- Shows a type switch to handle different types using `interface{}` and type assertions.

## 8. basics_08_arrays
Covers array fundamentals in Go, including declaration, access, modification, and iteration. The program:
- Declares and initializes a fixed-size array of integers.
- Accesses and modifies individual elements.
- Calculates array length and iterates using both traditional `for` loops and `range`.
- Demonstrates multi-dimensional arrays with nested loops for traversal.

## 9. basics_09_blank_identifiers
Introduces the blank identifier (`_`) in Go for ignoring unwanted values. The program:
- Uses `_` to discard one return value from a function while capturing the other.
- Ignores multiple return values entirely.
- Employs `_` in range loops to skip indices when only values are needed, such as summing a slice.

## 10. basics_10_slices
Explains slices in Go, which are dynamic arrays built on top of arrays. The program demonstrates:
- Declaring and initializing slices directly or from arrays.
- Slicing arrays to create sub-slices with start:end syntax.
- Modifying slices and observing how changes affect the underlying array (since slices share memory).
- Appending elements to slices using `append()`.
- Creating slices with `make()` and understanding length vs. capacity.

## 11. basics_11_maps
Covers maps (key-value data structures) in Go. The program shows:
- Declaring and initializing maps with literal syntax.
- Accessing values by key and handling missing keys.
- Adding and updating key-value pairs.
- Checking if a key exists using the comma ok idiom.
- Deleting keys with `delete()`.
- Iterating over maps with `range` (note: iteration order is not guaranteed).

## 12. basics_12_ranges
Demonstrates the `range` keyword in Go for iterating over collections. The program:
- Iterates over a string, showing both index and rune (Unicode code point) values for each character.
- Explains the difference between iterating over bytes vs. runes in strings (important for multi-byte Unicode).
- Shows `range` with arrays/slices (index and value) and maps (key and value).

## 13. basics_13_functions
Introduces functions in Go as first-class citizens. The program demonstrates:
- Defining functions with parameters and return types.
- Passing functions as arguments to other functions.
- Returning functions from other functions.
- Using function variables to store and call functions.
