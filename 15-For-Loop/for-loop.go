package main

import (
	"fmt"
)

func main() {
	// For loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("-----------------------")

	rows := 8 // Nested for loop
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows; j++ {
			if j <= rows-i {

				fmt.Print(" ")
			}
			fmt.Print("*")
		}
		fmt.Println()
	}

}

// For components
// 1. Initialization: This is executed once at the beginning of the loop.
// 2. Condition: This is checked before each iteration. If it evaluates to true, the loop continues; if false, the loop ends.
// 3. Post statement: This is executed after each iteration.
// 4. Break statement: This is used to exit the loop prematurely.
// 5. Continue statement: This is used to skip the current iteration and move to the next one.
// 6. Range: This is used to iterate over collections, such as arrays or slices.
// 7. Infinite loop: This is a loop that continues indefinitely until a break statement is encountered.
// 8. While loop: This is a loop that continues as long as a condition is true.
// 9. Traditional for loop: This is a loop that consists of three parts: initialization, condition, and post statement.

// For loop is a control flow statement that allows code to be executed repeatedly based on a condition. It consists of three parts: initialization, condition, and post statement. The initialization is executed once at the beginning of the loop, the condition is checked before each iteration, and the post statement is executed after each iteration. The loop continues until the condition evaluates to false.
// The break statement is used to exit the loop prematurely, while the continue statement is used to skip the current iteration and move to the next one. For loops can also be used to iterate over collections, such as arrays or slices, using the range keyword.
// In Go, the for loop is the only loop construct available, but it can be used in various ways to achieve different looping behaviors. It can be used as a traditional for loop, a while loop, or an infinite loop by omitting the condition and post statement.
