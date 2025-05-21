package main

import (
	"fmt"
)

// Custom error types can be created by implementing the error interface.
type MyError struct {
	Message string
	Code    int
}

// Implement the Error() method for MyError.
func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {

	err := &MyError{
		Message: "Something went wrong",
		Code:    404,
	}

	fmt.Println(err)
}

// *MyError means that err is a pointer to MyError.
// &MyError{} means that we are creating a new instance of MyError.
// The Error() method is called when we print the error.
// Custom errors can be used to provide more context about the error.
// This can be useful for debugging and logging.
