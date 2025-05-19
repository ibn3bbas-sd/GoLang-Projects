package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Color  string
	Width  int
	Height int
}
type Car struct {
	Make  string
	Model string
	Year  int
	Price float64
}

func main() {
	// Initializing a struct
	person1 := Person{
		Name:   "John",
		Age:    30,
		Color:  "Blue",
		Width:  70,
		Height: 180,
	}

	fmt.Println("Person 1:", person1)

	// Modifying a struct field
	person1.Age = 31
	fmt.Println("Updated Person 1 Age:", person1.Age)

	// Accessing struct fields
	fmt.Println("Person 1 Name:", person1.Name)
	fmt.Println("Person 1 Color:", person1.Color)
}

// Accessing and Modifying Structs: In Go, you can access and modify struct fields using the dot notation.
// You can also initialize a struct using a struct literal, which allows you to set the values of its fields at the time of creation.
