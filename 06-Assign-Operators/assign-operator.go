package main

import (
	"fmt"
)

func main() {

	a := 3
	fmt.Println("* Before: ", a)

	//a = a + 4
	a += 4
	fmt.Println("* After : ", a)
}
