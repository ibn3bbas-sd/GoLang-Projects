package main

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
