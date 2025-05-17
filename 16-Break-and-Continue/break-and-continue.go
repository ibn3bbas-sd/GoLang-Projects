package main

import (
	"fmt"
)

func main() {
	// Break statement
	fmt.Println("This is Break Statement:")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	println("")
	println("-----------")
	println("")

	// Continue statement
	fmt.Println("This is Continue Statement:")
	for j := 1; j <= 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}
}

// Break and continue statements are used to control the flow of loops in Go.
// The break statement is used to exit a loop prematurely, while the continue statement is used to skip the current iteration and move to the next one.
// In the first loop, the break statement is used to exit the loop when i equals 5.
// In the second loop, the continue statement is used to skip the iteration when j equals 5, so it will not print 5 but will continue with the next iteration.S
