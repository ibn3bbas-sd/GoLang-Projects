package main

import (
	"fmt"
	"time"
)

// Switch for Channel
// The switch statement is used to select a case based on the value of a variable.

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Hello from Channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "Hello from Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-channel1: // Receive message from channel1
			fmt.Println(message1)
		case message2 := <-channel2: // Receive message from channel2
			fmt.Println(message2)
		}
	}
}

/*
Switch satement for channel is used to select a case based on the value of a variable.
In this example, we have two channels, channel1 and channel2.
We start two goroutines that send messages to these channels after a delay.
The select statement is used to wait for messages from either channel.
When a message is received, it prints the message to the console.
*/
