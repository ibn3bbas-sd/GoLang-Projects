package main

import "fmt"

func main() {

	var operator string
	var num1, num2 float64

	fmt.Println("Hi this my Calculator")

	fmt.Print("Enter 1st Number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter Operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter 2nd Number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("%3f + %3f = %3f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%3f - %3f = %3f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%3f * %3f = %3f\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%3f / %3f = %3f\n", num1, num2, num1/num2)
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
		fmt.Println("Please try again.")

	}
}
