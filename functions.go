// In Go language, the parameters passed to a function are called actual parameters,
//  whereas the parameters received by a function are called formal parameters.
// Go program to illustrate the

// concept of the call by value: any changes made inside functions are not reflected in actual parameters of the caller.
package main

import "fmt"

// // function which swap values
// func swap(a, b int) int {

// 	var o int
// 	o = a
// 	a = b
// 	b = o

// 	return o
// }

// Main function
// func main() {
// 	var p int = 10
// 	var q int = 20
// 	fmt.Printf("p = %d and q = %d", p, q)

// 	// call by values
// 	swap(p, q)
// 	fmt.Printf("\np = %d and q = %d", p, q)
// }

//call by reference:any changes made inside the function are actually reflected in actual parameters of the caller.

// // function which swap values
func swap(a, b *int) int {
	var o int
	o = *a
	*a = *b
	*b = o

	return o
}

// Main function
func main() {

	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d", p, q)

	// call by reference
	swap(&p, &q)
	fmt.Printf("\np = %d and q = %d", p, q)
}
