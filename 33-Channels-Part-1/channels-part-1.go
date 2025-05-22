package main

// Channels are a way to communicate between goroutines.
// They are used to send and receive data between goroutines.
// A channel is created using the make function and can be used to send and receive values of a specific type.

func main() {
	// Create a channel of type int
	channel := make(chan string)
	// Start a goroutine to send data to the channel
	go func() {
		for i := 0; i < 5; i++ {
			channel <- "Hello from goroutine!"
		}
		close(channel) // Close the channel when done
	}()
	// Receive data from the channel in the main function
	for message := range channel {
		println(message)
	}

}
