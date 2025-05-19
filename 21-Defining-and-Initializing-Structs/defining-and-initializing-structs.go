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

	// Initializing a struct with default values
	person2 := Person{
		Name:   "Jane",
		Age:    25,
		Color:  "Red",
		Width:  65,
		Height: 170,
	}

	fmt.Println("Person 2:", person2)

	// Initializing a struct with all fields
	car1 := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2020,
		Price: 24000,
	}

	fmt.Println("Car 1:", car1)

	// Initializing a struct with some fields
	car2 := Car{
		Make:  "Honda",
		Model: "Civic",
		Year:  2021,
		Price: 22000,
	}
	fmt.Println("Car 2:", car2)
}

// Defining and initializing structs in Go is straightforward. You can define a struct using the `type` keyword followed by the struct name and its fields.
// To initialize a struct, you can use a composite literal, which allows you to specify the values for each field in the struct.
// You can also use named fields to make your code more readable.
// In this example, we defined two structs: `Person` and `Car`, and initialized them with different values.
