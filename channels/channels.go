package main

import (
	"fmt"
)

/* 
 Channels are pipes that you can use to pass values from one goroutine to another
 Channels allow goroutines to share memory bu communicating
 We use the channel operator, <-, to send and receive values (data flows in the direction of the arrow)
 We use the `make` builtin function and the `chan` keyword + the type we will be sending 
 ch := make(chan type)
*/

func writeMessageToChannel(message chan string) {
	message <- "Hello you!"
}

func main() {
	fmt.Println("Channel Demo")
	// Create a channel
	message := make(chan string)
	// Call the function that returns the message value
	go writeMessageToChannel(message)
	
	// The program is blocked until we receive a response from the channel
	fmt.Println("Greeting from the channel: ", <-message)
	// Once we are done using the channel we close it
	close(message)
}