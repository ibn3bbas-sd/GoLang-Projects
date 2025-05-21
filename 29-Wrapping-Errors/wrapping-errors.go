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

	// The %w verb is used to wrap the error with additional context.
	// errors.New creates a new error with the message "can not divide by zero"
	// fmt.Errorf creates a new error with the message "division by zero: can not divide by zero"

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
	// Unwrapping errors is a way to retrieve the original error from a wrapped error.
	// This is done by using the errors.Unwrap function.
	// the errors.Is function checks if the error is of a specific type.

	unwrappedErr := errors.Unwrap(err)
	if unwrappedErr != nil {
		fmt.Println("Unwrapped Error:", unwrappedErr)
	}
	if errors.Is(err, unwrappedErr) {
		fmt.Println("The error is the same as the unwrapped error")
	} else {
		fmt.Println("The error is not the same as the unwrapped error")
	}
}
