package main

import "fmt"

// Nested Structs: In Go, you can define a struct within another struct.
// This is called a nested struct.
// It allows you to create complex data structures that can represent real-world entities more accurately.

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}
type Person struct {
	Name    string
	Age     int
	Address Address // Nested struct
}
type Car struct {
	Make  string
	Model string
	Year  int
	Price float64
	Owner Person // Nested struct
}

func main() {
	// Initializing a nested struct
	address1 := Address{
		Street:  "123 Main St",
		City:    "Makkah",
		State:   "Jeddah",
		ZipCode: "10001",
	}
	person1 := Person{
		Name:    "Abdullah",
		Age:     30,
		Address: address1,
	}
	car1 := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2020,
		Price: 24000.00,
		Owner: person1,
	}

	fmt.Println("Car 1:", car1)
	fmt.Println("Car Owner Name:", car1.Owner.Name)
	fmt.Println("Car Owner Address:", car1.Owner.Address.Street, car1.Owner.Address.City)
}
