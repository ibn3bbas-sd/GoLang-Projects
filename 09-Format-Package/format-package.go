package main

// The format package is used to format strings, numbers, and other data types.
// It provides functions to format data in a variety of ways, including
//fmt means "format" and is used to format strings, numbers, and other data types.
// \n is used to create a new line in the output.
// %d is used to format integers.
// %f is used to format floating-point numbers.
// %s is used to format strings.
// %v is used to format any value.

//fmt.scanln is used to read input from the user.

import (
	"fmt"
)

func main() {

	a := 12    //int
	b := 34.56 //float64

	fmt.Printf("int: %v\n", a)     //int
	fmt.Printf("float: %.3v\n", b) //float64

	var name string

	fmt.Scanln(&name) //input from user

}
