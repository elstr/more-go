package main

import "fmt"

func main() {
 messageQueue := make(chan string, 3)

 messageQueue <- "first"
 messageQueue <- "second"
 messageQueue <- "third"

 // We can close the channel but since it's non-empty, we can still
 // receive the remaining values 
 close(messageQueue)

 // We use the range keyword to iterate over each element as it gets
 // received from the messageQueue
 for m := range messageQueue {
     fmt.Println(m)
 }
 
}