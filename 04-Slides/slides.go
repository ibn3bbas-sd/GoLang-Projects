package main

//Slides concept
// 1. Slice is a reference type
// 2. Slice is a dynamic array
// 3. Slice is a pointer to an array
// 4. Slice is a data structure that describes a section of an array
// 5. Slice is a flexible and powerful data structure

import "fmt"

func main() {

	var numbers []int
	names := []string{"Abbas", "Ali", "Zahra", "Fatemeh"}
	names = append(names, "Mohammad", "Reza", "Sara", "Niloofar")

	namesLength := len(names)

	fmt.Println(numbers)
	fmt.Println(names)
	fmt.Println(namesLength)
}
