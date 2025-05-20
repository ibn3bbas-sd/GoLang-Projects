package main

import (
	"errors"
	"fmt"
)

// Error handling in Go is done using the built-in error type.
// The error type is an interface that has a single method Error() string.
// When a function returns an error, it typically returns a value of type error.
// You can check if an error is nil to determine if the function succeeded or failed.
// You can also use the errors package to create and manipulate errors.

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
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
