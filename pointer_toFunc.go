// Go program to create a pointer
// and passing the address of the
// variable to the function
package main

import "fmt"

// taking a function with integer
// type pointer as an parameter
func ptf(a *int) {

	// dereferencing
	*a = 456
}

// Main function
func main() {

	// taking a normal variable
	var b = 100

	fmt.Printf("The value of x before function call is: %d\n", b)

	// calling the function by
	// passing the address of
	// the variable x
	ptf(&b)

	fmt.Printf("The value of x after function call is: %d\n", b)

	//returning pointer from a function
	n := rpf()

	// displaying the value
	fmt.Println("Value of n is: ", *n)
}

// defining function having integer
// pointer as return type
func rpf() *int {

	// taking a local variable
	// inside the function
	// using short declaration
	// operator
	lv := 200

	// returning the address of lv
	return &lv

}
