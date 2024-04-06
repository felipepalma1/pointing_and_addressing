package main

import "fmt"

func main() {

	//How to store addresses.
	//Pointers are variables that specifically store addresses.

	//We even set the data type of the addresses value for the pointer. For instance
	//var pointerForInt *int

	//In this example, pointerForInt will store the address of a variable that has an int data type.
	//To break it down further, the * operator signifies that this variable will store an address and the int portion
	// means that the address contains an integer value

	//With pointerForInt initialized, we can assign it value like so

	//var pointerForInt *
	//minutes := 525600
	//pointerForInt = &minutes
	//fmt.Println(pointerForInt) // Prints 0xc000018038

	star := "Polaris"
	var starAddress *string = &star
	fmt.Println("The address of star is", starAddress)
}
