package main

import (
	"fmt"
)

func main() {

	month := 7 // Example month
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Invalid month")
	}
}

// Switch case is a control structure that allows you to execute different blocks of code based on the value of a variable or expression. It is often used as an alternative to multiple if-else statements when you have many conditions to check.
// In the example above, we use a switch statement to check the value of the variable month. Depending on the value, it prints the corresponding month name. If the value doesn't match any case, it executes the default case and prints "Invalid month".
// The switch statement is a cleaner and more readable way to handle multiple conditions compared to using multiple if-else statements. It also allows for fall-through behavior, where multiple cases can execute the same block of code if not explicitly handled.
// In Go, the switch statement is a powerful and flexible control structure that can be used to simplify complex conditional logic. It allows you to evaluate an expression and execute different blocks of code based on the value of that expression. The switch statement can be used with various types of expressions, including integers, strings, and even boolean values.
// The switch statement can also be used without an expression, in which case it acts like a series of if-else statements. This allows you to write more complex conditions and handle multiple cases in a single switch statement.
