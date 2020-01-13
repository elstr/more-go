package main

import (
	"fmt"
	"time"
)

// GORoutines are concurrent functions. 
// When the main func ends the routine is ended as well even if it didn't finish.
// Check the channels folder to see how to fix this :)

func printGreeting(source string){
	for i:=0; i < 9; i++ {
		fmt.Println("Hello!", i, source)
	}
	// this isn't a cool solution
	// we don't know how much time the routine needs to finish
	time.Sleep(time.Millisecond * 2)
}

func main() {
	go printGreeting("routine")
	printGreeting("main")
}