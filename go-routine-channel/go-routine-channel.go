package main

import (
	"fmt"
)

var done chan bool = make(chan bool)

func printGreeting(source string){
	for i:=0; i < 9; i++ {
		fmt.Println("Hello!", i, source)
	}
	if source == "go-routine" {
		done <- true
	}
}

func main() {

	go printGreeting("go-routine")
	printGreeting("main")

	<-done
	close(done)
}