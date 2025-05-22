package main

import (
	"sync"
	"time"
)

var (
	counter int
	mutex   sync.Mutex
)

// Mutex to protect the counter
func increment() {
	for i := 0; i < 1500; i++ {
		mutex.Lock()   // Lock the mutex before accessing the counter
		counter++      // Increment the counter
		mutex.Unlock() // Unlock the mutex after accessing the counter
	}
}
func main() {
	go increment()
	go increment()

	time.Sleep(1 * time.Second)

	println("Final counter value:", counter)

}
