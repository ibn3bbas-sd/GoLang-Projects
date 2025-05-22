package main

import "fmt"

// Receive Messages from a Channel
// Send Messages to a Channel

func sendMessages(channel chan<- string, message string) {
	// Send messages to the channel
	channel <- message
}

func receiveMessages(channel <-chan string) string {
	// Receive messages from the channel
	return <-channel

}

func main() {
	// Create a channel of type string
	channel := make(chan string)

	// Start a goroutine to send messages to the channel
	go sendMessages(channel, "Hello from goroutine!")

	// Receive messages from the channel in the main function
	message := receiveMessages(channel)
	fmt.Println(message)
}
