package main

import "fmt"

// Variadic function that takes a variable number of arguments

func classNames(class ...string) {
	// Loop through the class names and print them
	for i, name := range class {
		fmt.Printf("Class %d: %s\n", i+1, name)
	}
}

func main() {
	// Call the variadic function with different number of arguments
	classNames("Math", "Science", "History")
	classNames("English", "Art")
	classNames("PE")
}

// %d is a format specifier for integers
// %s is a format specifier for strings
// /n is a newline character

// Variadic functions are useful when you don't know how many arguments will be passed to the function.
// They allow you to pass a variable number of arguments to a function.
// In this example, the classNames function takes a variable number of string arguments.
// The ... operator indicates that the function can take a variable number of arguments.
// The function then loops through the arguments and prints them out.
// Variadic functions can be used in many different scenarios, such as:
// - Logging messages with variable number of parameters
// - Handling events with variable number of parameters
// - Creating flexible APIs that can accept different types of input
// - Building command-line tools that accept variable number of arguments
// - Creating flexible data structures that can hold variable number of elements
