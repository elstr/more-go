package main

import (
	"fmt"
	"os"
	"os/signal" // Identify when an interruption occures
	"time"
)

// Used to print the values always in the same terminal spot
// Score board first slot will show PI value result
// Score board second slot will show number of Nilakantha series we've processed
const ANSIClearScreenSequence = "\033[H\033[2J"
const ANSIFirstSlotScreenSequence = "\033[2;0H"
const ANSISecondSlotScreenSequence = "\033[3;0H"
const ANSIThirdSlotScreenSequence = "\033[4;0H"

// Channel used to update the value of pi on the scoreboard
var pichan chan float64 = make(chan float64)
// Channel used to indicate that the program can finish
var finished chan bool = make(chan bool)

var ch chan float64 = make(chan float64)

// Number of Nilakantha terms for the scoreboard
var termsCount int

// This function renders our virtual scoreboard
// with current computed value of Pi using Nilakantha's formula
func printCalculationSummary() {
	fmt.Print(ANSIClearScreenSequence)
	fmt.Println(ANSIFirstSlotScreenSequence, "Computed value of PI: \t\t", <-pichan)
	fmt.Println(ANSISecondSlotScreenSequence, "Nilakantha terms: \t\t", termsCount)
	fmt.Println(ANSISecondSlotScreenSequence, "Chan value: \t\t", <-ch)
}

func pi(n int) float64 {
	// ch := make(chan float64)

	for k := 1; k <= n; k++ {
		// Each Nilakantha term is calculated in its own goroutine
		go nilakanthaTerm(ch, float64(k))
	}

	f := 3.0

	// We sum up the calculated Nilakantha terms for n steps
	for k := 1; k <= n; k++ {
		termsCount++
		f += <- ch
		pichan <- f
	}

	finished <- true
	return f

}

func nilakanthaTerm(ch chan float64, k float64){
	j := 2 * k
	if int64(k)%2 == 1 {
		ch <- 4.0 / (j * (j+1) * (j+2))
	} else {
		ch <- -4.0 / (j * (j+1) * (j+2))
	}
}


func main() {

	ticker := time.NewTicker(time.Millisecond * 400)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go pi(5)

	go func() {
		for range ticker.C {
			printCalculationSummary()
		}
	}()

	for {
		select {
			case <-finished:
				ticker.Stop()
				fmt.Println("Program done calculating Pi.")
				os.Exit(0)
		
			case <-interrupt:
				ticker.Stop()
				fmt.Println("Program interrupted by the user.")
				return
		}

	}
}