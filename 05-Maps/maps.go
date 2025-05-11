package main

import "fmt"

// Maps are a built-in data type in Go that allows you to store key-value pairs.

// Maps concept:
// 1. Maps are unordered collections of key-value pairs
// 2. Maps are reference types
// 3. Maps are dynamic data structures
// 4. Maps are implemented as hash tables

func main() {
	var ages = map[string]int{
		"Ahmed":    25,
		"Ali":      30,
		"Zahra":    28,
		"Mohammad": 35,
	}
	fmt.Println(ages["Ali"])
	fmt.Println(ages["Zahra"])
	fmt.Println(ages["Mohammad"])
	fmt.Println(ages["Ahmed"])

}
