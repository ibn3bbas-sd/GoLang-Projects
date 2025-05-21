package main

import (
	"log"
	"os"
)

func main() {
	data := []byte("Hello, World!\n")
	err := os.WriteFile("hello.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File written successfully")
}
