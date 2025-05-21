package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("secrets.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
