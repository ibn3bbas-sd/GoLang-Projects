package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func areEqual(p1, p2 Person) bool {
	return p1.Name == p2.Name && p1.Age == p2.Age && reflect.DeepEqual(p1.Hobbies, p2.Hobbies)
}

func main() {
	person1 := Person{Name: "Sawda", Age: 3, Hobbies: []string{"Playing", "Chatting"}}
	person2 := Person{Name: "Ihsan", Age: 28, Hobbies: []string{"Reading", "Traveling"}}

	if areEqual(person1, person2) {
		fmt.Println("person1 and person2 are equal")
	} else {
		fmt.Println("person1 and person2 are not equal")
	}

	person3 := person1
	fmt.Println("person3:", person3)
}

// Comparison of structs: In Go, you can compare structs directly using the == operator.
// This compares each field of the struct for equality.
// However, if the struct contains fields that are not comparable (like slices or maps), you will get a compilation error.
// To compare structs with non-comparable fields, you can use the reflect package to perform a deep comparison.
