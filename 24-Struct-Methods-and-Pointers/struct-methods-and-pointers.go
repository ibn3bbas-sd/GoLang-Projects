package main

import "fmt"

// Struct Methods: In Go, you can define methods on structs.
// This allows you to associate behavior with the data structure.
// Methods are functions that have a receiver argument.
// The receiver is a special parameter that allows you to call the method on an instance of the struct.

type Rectangle struct {
	Length float64
	Width  float64
}
type Circle struct {
	Radius float64
}
type Square struct {
	Side float64
}
type Triangle struct {
	Base   float64
	Height float64
}

// Method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Method to calculate the area of a circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Method to calculate the area of a square
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Method to calculate the area of a triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
func main() {
	// Create instances of the structs
	rect := Rectangle{Length: 5, Width: 3}
	circ := Circle{Radius: 4}
	sq := Square{Side: 2}
	tri := Triangle{Base: 6, Height: 4}

	// Call the Area method on each instance
	fmt.Println("Area of Rectangle:", rect.Area())
	fmt.Println("Area of Circle:", circ.Area())
	fmt.Println("Area of Square:", sq.Area())
	fmt.Println("Area of Triangle:", tri.Area())
}

// Path by value vs by reference: In Go, when you pass a struct to a function or method, it is passed by value.
// This means that a copy of the struct is made, and any changes made to the copy do not affect the original struct.
// If you want to modify the original struct, you can pass a pointer to the struct instead.
// To do this, you can use the * operator to define a pointer receiver in the method definition.
