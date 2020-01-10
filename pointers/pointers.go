package main

/*
	Pointers analogy:
	You have a home address and a key to get inside. 
	Pointer work in the same way.
	A pointer is the address to your home and the key to access it.
	It points to a certain memory address and gives you access to it so you get the value/s stored. 
*/

import (
	"fmt"
)

func main() {
	var myAddress string = "White House 123"
	var myAddressPointer *string // the * operator says that string is a pointer

	myAddressPointer = &myAddress // the & get's the memory address of the string variable
	
	fmt.Println(myAddressPointer) // returns 0xc0000581e0 the address of myAddress string variable

	// OK so we got home... but how do we actually get inside ?? How do we get the value of the memory address our pointer points to
	fmt.Println(*myAddressPointer) // using the * before the pointer says to go HEY! Get me this pointer's value pls

	// OK we got home and we got inside :) Yay! 

	// What use can we give pointers ? 
	// GO is a "pass-value" language that means when we pass a parameter we get a copy of it and modify the copy not the real var
	// So let's say you want to double a number by calling a function:
	num := 10 // := the shorthand way to create an integer and assign value
	doubleThisNumber(num)
	fmt.Println(num) // prints out 10 !! and we wanted 10 doubled :(
	// Pointer to the rescue but we need to change our function to receive a pointer now (let's create a new function)
	num2 := 50
	// We pass the ADDRESS of num2 because doubleWithPointer expects a pointer
	// doubleWithPointer will modify the VALUE of num2 using the asterisc operator (check the doubleWithPointer func line with *theNumberPointer)
	doubleWithPointer(&num2) 
	fmt.Println(num2) // The value of the variable has been changed 
}

func doubleThisNumber(theNumber int) int {
	return theNumber * 2
}

// doubleWithPointer receives a pointer and returns a pointer 
// that's why there are 2 asteriscs *: (theNumberPointer *int) *int
// we could omit the pointer returned but this is just an example given
func doubleWithPointer(theNumberPointer *int) *int { 
	*theNumberPointer *= 2 // remember that by using the * before the pointer we access the pointer's value
	return theNumberPointer // we return the pointer here
}