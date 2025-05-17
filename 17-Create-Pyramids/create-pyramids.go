package main

import "fmt"

func main() {
	rows := 9 // Number of rows for the pyramid

	for i := 1; i <= rows; i++ {
		// Print spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// Print stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	println("")
	println("-----------")
	println("")

	newRows := 9 // Number of rows for the pyramid

	for i := 1; i <= newRows; i++ {
		// Print spaces
		for j := 1; j <= newRows-i; j++ {
			fmt.Print(" ")
		}
		// Print numbers
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print(i)
		}
		fmt.Println()
	}

}

// Explain the code
// The code above creates two pyramids using nested loops.
// The first pyramid is made of stars (*) and the second pyramid is made of numbers.
// The outer loop iterates through the number of rows, while the inner loops handle the spaces and the characters to be printed.
// The first pyramid is centered by printing spaces before the stars, and the second pyramid is also centered by printing spaces before the numbers.
// The number of stars or numbers printed in each row is determined by the formula 2*i-1, where i is the current row number.
// The first pyramid has stars, while the second pyramid has the row number printed multiple times.
// The code uses the fmt package for printing to the console.
// The first pyramid is a classic pyramid shape, while the second pyramid has a unique pattern where each row contains the same number repeated.
// The pyramids are printed to the console with a blank line and a separator line in between.
