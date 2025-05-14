package main

// If statement is used to execute a block of code if a specified condition is true.
// If the condition is false, the code block will not be executed.
// The if statement can be used with else and else if statements to create multiple conditions.
// The if statement can also be used with logical operators (&&, ||) to combine multiple conditions.
// The if statement can be used with comparison operators (==, !=, >, <, >=, <=) to compare values.

import (
	"fmt"
)

func main() {
	experience := 10
	certification := "Microsoft certified: Azure DevOps Engineer Expert"

	if experience <= 5 && certification == "Microsoft certified: Azure DevOps Engineer Expert" {
		fmt.Println("You are a Senior DevOps Engineer")
	} else {
		fmt.Println("You are a Junior DevOps Engineer")
		fmt.Println("You are a Mid-level DevOps Engineer")
	}
}
