package main

import "fmt"

func add(x int, y int) int {
	return x + y
	// Return a single value
}

func divide(x int, y int) (int, int) {

	quotient := x / y
	remainder := x % y
	return quotient, remainder
	// Return multiple values
}

func calculate(x int, y int) (sum int, product int) {
	// Named return values
	sum = x + y
	product = x * y
	return
	// Implicit return
}

func main() {
	// Function with multiple parameters
	result := add(2, 3)
	fmt.Println("The result is:", result)

	q, r := divide(25, 3)
	fmt.Println("The quotient is:", q)
	fmt.Println("The remainder is:", r)
	// Function call with different arguments

	s, p := calculate(4, 6)
	fmt.Println("The sum is:", s)
	fmt.Println("The product is:", p)
	// Function call with different arguments
}
