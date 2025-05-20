package main

import "fmt"

// Interface in Go is a type that specifies a contract by defining methods that a type must implement.
// It allows for polymorphism (تعدد الاشكال), where different types can be treated as the same type if they implement the same interface.
// An interface is defined using the `interface` keyword, followed by a set of method signatures.

type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}
func (d Dog) Move() string {
	return "I run on four legs."
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow! My name is " + c.Name
}
func (c Cat) Move() string {
	return "I walk gracefully."
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Create a slice of Animal interface
	animals := []Animal{dog, cat}

	// Iterate over the slice and call methods
	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
	}
}
