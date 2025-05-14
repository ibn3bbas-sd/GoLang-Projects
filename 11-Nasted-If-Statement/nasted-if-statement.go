package main

import (
	"fmt"
)

func main() {

	if number := 10; number > 0 {
		if number%2 == 0 {
			fmt.Println("The number is positive and even.")
		} else {
			fmt.Println("The number is positive and odd.")
		}
	}
	if number := -5; number < 0 {
		if number%2 == 0 {
			fmt.Println("The number is negative and even.")
		} else {
			fmt.Println("The number is negative and odd.")
		}
	}
}
