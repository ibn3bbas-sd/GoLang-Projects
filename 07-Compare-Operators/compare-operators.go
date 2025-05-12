package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	fmt.Println("a == b:", a == b) // false
	fmt.Println("a != b:", a != b) // true
	fmt.Println("a > b :", a > b)  // false
	fmt.Println("a < b :", a < b)  // true
	fmt.Println("a >= b:", a >= b) // false
	fmt.Println("a <= b:", a <= b) // true
}
