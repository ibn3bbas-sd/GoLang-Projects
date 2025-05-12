package main

// Logical Operators
// Logical operators are used to combine multiple boolean expressions.
// The logical operators in Go are:
// 1. AND (&&)
// 2. OR (||)
// 3. NOT (!)
// The AND operator returns true if both operands are true.
// The OR operator returns true if at least one of the operands is true.
// The NOT operator negates the boolean value of the operand.

import (
	"fmt"
)

func main() {

	num1 := 10
	num2 := 20
	num3 := 30

	fmt.Println(num1 < num2 && num2 < num3) // true
	fmt.Println(num1 < num2 && num2 > num3) // false

	fmt.Println(num1 < num2 || num2 < num3) // true
	fmt.Println(num1 > num2 || num2 < num3) // true

	fmt.Println(!(num1 > num2 || num2 < num3)) // false
	fmt.Println(!(num1 < num2 && num2 < num3)) // false

}
