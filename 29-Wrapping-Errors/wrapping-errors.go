package main

import (
	"errors"
	"fmt"
)

// Wrapping errors is a way to add context to an error.
// This is done by using the fmt.Errorf function with the %w verb.
// The %w verb allows you to wrap an error with additional context.

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero: %w", errors.New("can not divide by zero"))
	}
	return a / b, nil
}
func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
