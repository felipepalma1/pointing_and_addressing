package main

import "fmt"

// We know that addresses are where values are stored and pointers allow us to keep track of addresses.
// But what if we want the address to store a different value? Well we can actually use our pointer to access
// the address and change its value. This action is called dereferencing or indirecting

//We need to use the * operator again on a pointer and then assign a new value like so

func main() {
	star := "Polaris"
	starAddress := &star
	*starAddress = "Sirius"
	fmt.Println("The actual value of star is", star)
}
