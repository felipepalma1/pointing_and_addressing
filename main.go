package main

import "fmt"

func addHundred(num int) {
	num += 100
}

// Go is a pass-by-value language. We are passing functions the value of an argument.
// In a technical sense, when we are calling a function with an argument, the Go compiler is strictly using the value
// of the argument rather than the argument itself. Because of this feature (pass-by-value), the changes that take place
// in our function, stay within that function.

// We can change the values of different scopes. To do so, we need to make use of
// * Addresses
// * Pointers
// * Dereferencing

//
// En este ejemplo se envia el valor de x a la funcion. Se usa el valor para sumarle 100, pero no se altera
// la variable x. Solo se utiliza el numero que tiene asignado.

func main() {

	x := 1
	addHundred(x)
	fmt.Println(x)
}
