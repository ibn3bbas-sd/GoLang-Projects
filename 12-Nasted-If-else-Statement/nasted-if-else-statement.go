package main

import (
	"fmt"
)

func main() {
	day := 1 // Assigning a value to day

	if day == 1 {
		fmt.Println("Today is Saturday")
	} else if day == 2 {
		fmt.Println("Today is Sunday")
	} else if day == 3 {
		fmt.Println("Today is Monday")
	} else if day == 4 {
		fmt.Println("Today is Tuesday")
	} else if day == 5 {
		fmt.Println("Today is Wednesday")
	} else if day == 6 {
		fmt.Println("Today is Thursday")
	} else if day == 7 {
		fmt.Println("Today is Friday")
	} else {
		fmt.Println("Invalid day")
	}
}

// In this code, we are using a nested if-else statement to check the value of the variable 'day'.
// Depending on the value of 'day', it will print the corresponding day of the week.
// If the value of 'day' is not between 1 and 7, it will print "Invalid day".
