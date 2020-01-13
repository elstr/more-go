package main

import "fmt"

func main() {
	// Create a buffer channel with capacity of 3 and type string
	messageQueue := make(chan string, 3)
	// Pass values to our buffer
	messageQueue <- "first"
	messageQueue <- "second"
	messageQueue <- "third"

	// Drain the buffer passing the 3 values
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
}