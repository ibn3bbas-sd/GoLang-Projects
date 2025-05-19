package main

import "fmt"

func main() {

	// Declare a variable of type int
	var x int = 10

	// Declare a pointer variable that points to the address of x
	var p *int = &x

	// Print the value of x and the address of x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)

	// Print the value of p and the address of p
	fmt.Println("Value of p:", p)
	fmt.Println("Address of p:", &p)

	// Print the value pointed to by p
	fmt.Println("Value pointed to by p:", *p)

	// Change the value of x using the pointer p
	*p = 20

	// Print the new value of x
	fmt.Println("New value of x:", x)
}

// Pointers are variables that store the memory address of another variable.
// The * operator is used to declare a pointer variable and to dereference a pointer.
// The & operator is used to get the address of a variable.
// In this example, we declare a variable x of type int and a pointer variable p that points to the address of x.
// We then print the value of x, the address of x, the value of p, the address of p, and the value pointed to by p.
// Finally, we change the value of x using the pointer p and print the new value of x.
