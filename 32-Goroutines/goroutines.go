package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent tasks in Go programs.
// A goroutine is created by using the go keyword followed by a function call.
// The function will run concurrently with the rest of the program.

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// Start a goroutine
	go printMessage("Hello from goroutine!")

	// Main function will also run concurrently
	printMessage("Hello from main function!")

	// Wait for the goroutine to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}
