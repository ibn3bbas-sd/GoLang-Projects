package main

import (
	"fmt"
	"math"
)

func add(x int, y int, z int) int {
	return x + y + z
}

func maltiply(x int, y int, z int) int {
	return x * y * z
}

func power(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func main() {

	result := add(2, 3, 4)
	fmt.Println("The sum is:", result)
	// Function call

	product := maltiply(2, 3, 4)
	fmt.Println("The product is:", product)
	// Function call with different arguments

	end := power(2, 34)
	fmt.Println("The power is:", end)
	// Function call with different arguments
}

// Function with multiple parameters
// Forbidden to use the same name for different parameters
// Function with different types of parameters
