
# learning-golang

Each folder contains a `main.go` file with example code and comments for easy understanding.

# basics concepts

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
- Functions returning multiple values of the same type.
- Functions returning multiple values of different types.
- Functions returning a value along with an error for error handling.


## 14. basics_14_defer
Explains the `defer` keyword in Go, which schedules function calls to be executed after the surrounding function returns. The program demonstrates:
- Using `defer` for resource cleanup (e.g., closing files or network connections).
- Defer execution order, which follows Last In, First Out (LIFO) principle.
- How `defer` interacts with variables - the value is evaluated at the time the defer statement is executed, not when the deferred function runs.


## 15. basics_15_panic
Introduces the `panic` function in Go, which stops normal execution and begins panicking. The program demonstrates:
- Triggering a panic with a custom message when a condition is met.
- How panic causes the program to crash with a stack trace.
- Using `defer` to execute cleanup code before the panic unwinds the stack.


## 16. basics_16_recover
Explains the `recover` function in Go, used to regain control of a panicking goroutine. The program demonstrates:
- Using `recover()` inside a deferred function to catch panics.
- Preventing the program from crashing by handling the panic gracefully.
- Printing the recovered panic value for debugging purposes.

## 17. basics_17_exit
Demonstrates program termination using `os.Exit()`. The program shows:
- Using `os.Exit()` to terminate the program immediately with an exit code.
- How `os.Exit()` bypasses deferred function calls (unlike normal program termination).
- The difference between normal return and `os.Exit()` in terms of cleanup execution.

## 18. basics_18_init
Explains the `init()` function in Go, which runs automatically before `main()`. The program demonstrates:
- Declaring and using init functions for setup and initialization.
- How multiple init functions are executed in sequential order.
- Init functions running before the main function, useful for configuring settings or initializing variables.

# Intermediate Concepts

## 1. int_01_closures
Introduces closures in Go, which are functions that capture and remember variables from their surrounding lexical scope, even after the outer function has finished executing. The program demonstrates:
- Creating a closure with the `add()` function that maintains state between calls, accumulating a sum.
- How each closure instance has its own independent state, demonstrated by creating multiple adder functions.
- Using anonymous functions to create closures, such as a `multiplier` function that returns a closure for multiplying by a specific factor.
- Practical examples showing how closures can be used to create functions with persistent state without global variables.
- Creating specialized functions like `double` and `triple` that remember their multiplication factors.

## 2. int_02_recursion
Introduces recursion in Go, where a function calls itself to solve problems. The program demonstrates generating the Fibonacci sequence recursively, showing how each call builds upon the previous ones to compute the sequence up to a given number of terms.

## 3. int_03_pointers
Introduces pointers in Go, which store memory addresses of variables. The program demonstrates:
- Declaring pointer variables using `*type` syntax.
- Getting the memory address of a variable using the `&` operator.
- Dereferencing pointers to access the value they point to using `*pointer`.
- Modifying the original variable's value through the pointer.
- Passing pointers to functions to allow the function to modify the caller's variables.

## 4. int_04_strings_runes
Introduces strings and runes in Go, which are essential for handling text and Unicode characters. The program demonstrates:
- Iterating over strings using `range` to access both index and rune values.
- The difference between escape sequences in double quotes vs. backticks.
- String comparison using lexicographical order.
- Counting characters in Unicode strings using `utf8.RuneCountInString`.
- String concatenation using the `+` operator.
- Using runes to represent individual Unicode characters.

## 5. int_05_formatters
Explains various format specifiers in Go's `fmt` package for controlling output formatting. The program demonstrates:
- General format verbs like `%v`, `%#v`, `%T`, and `%%`.
- Integer formatting in different bases: decimal (`%d`), binary (`%b`), hexadecimal (`%x`, `%X`), and octal (`%o`).
- String formatting with width, alignment, and truncation using `%s`, `%q`.
- Float formatting with precision and scientific notation using `%f`, `%g`, `%e`, `%E`.
- Boolean formatting using `%t`.

## 6. int_06_structs
Covers structs in Go, which are composite data types that group related fields. The program demonstrates:
- Defining structs with named fields and anonymous fields.
- Embedding structs to create composite types.
- Creating struct instances using struct literals and field assignment.
- Defining methods on structs with value receivers and pointer receivers.
- Using anonymous structs for one-off data structures.

## 7. int_07_interfaces
Introduces interfaces in Go, which define sets of methods for polymorphism and abstraction. The program demonstrates:
- Defining interfaces with method signatures.
- Implementing interfaces by providing methods on structs.
- Using interfaces as function parameters for polymorphic behavior.
- How structs can implement multiple interfaces.
- Calling methods through interface values.
- Compile-time checks ensuring structs fully implement interfaces (demonstrated with a manufacturer struct that only partially implements the interface).

## 8. int_08_generics
Introduces generics in Go (available in Go 1.18+), which allow writing flexible, type-safe code that works with multiple types. The program demonstrates:
- Defining generic structs using type parameters (e.g., `Stack[T any]`).
- Implementing methods on generic types with receivers and type parameters.
- Creating generic stack implementations that work with different types (int, string, etc.).
- The `push()` method to add elements to the stack using `append()`.
- The `pop()` method to remove and return elements, handling the empty stack case by returning a zero value and false.
- The `isEmpty()` method to check if the stack has elements.
- Practical usage by instantiating and using the same generic stack type with different concrete types.

## 9. int_09_errors
Introduces error handling in Go, which is a fundamental part of writing robust programs. The program demonstrates:
- Creating custom errors using `errors.New()`.
- Functions returning both a result and an error.
- Checking for errors after function calls and handling them appropriately.
- Printing error messages to inform users about what went wrong.
- Using conditional logic to branch based on whether an operation succeeded or failed.

## 10. int_10_customerrors
Introduces custom error types in Go, allowing for more structured and informative error handling. The program demonstrates:
- Defining custom error structs that implement the error interface.
- Creating custom error constructors for consistent error creation.
- Using `errors.As()` for type assertion to check and extract specific error types.
- Handling different types of errors with appropriate responses.
- Simulating real-world scenarios like HTTP errors with status codes and messages.

## 11. int_11_stringfunc
Introduces string manipulation functions in Go, showcasing the `strings` package and related utilities. The program demonstrates:
- Converting strings to uppercase and lowercase using `strings.ToUpper` and `strings.ToLower`.
- Checking for substrings with `strings.Contains`, replacing with `strings.Replace`, and splitting/joining with `strings.Split` and `strings.Join`.
- Trimming whitespace with `strings.TrimSpace`, checking prefixes/suffixes with `HasPrefix` and `HasSuffix`.
- Finding indices with `strings.Index`, substring extraction via slicing.
- String concatenation, repetition with `strings.Repeat`, and length checks.
- Working with ASCII and Unicode values, byte vs. rune counts.
- Simple email validation, integer to string conversion with `strconv.Itoa`.
- Counting character occurrences, regular expressions with `regexp` for digit extraction.
- Handling Unicode strings with `utf8.RuneCountInString`.
- Efficient string building with `strings.Builder` for concatenation.

## 12. int_12_stringformat
Introduces string formatting in Go using the `fmt` package. The program demonstrates:
- Using `fmt.Sprintf` to format strings and store the result in a variable.
- Using `fmt.Printf` to print formatted output directly to the console.
- Formatting numbers with leading zeros using width specifiers like `%05d`.
- Formatting strings with width and alignment using `%10s` (right-aligned) and `%-10s` (left-aligned).
- The difference between interpreted string literals (`"Hello \nWorld"`) and raw string literals (`` `Hello \nWorld` ``) for preserving escape sequences.