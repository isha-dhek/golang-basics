// Golang program to demonstrate the declaration
// and initialization of pointers
package main

import "fmt"

func main() {

	// taking a normal variable
	var x int = 7334

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &x

	// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", *p)

	//double pointer:When we define a pointer to pointer then the first pointer is used to store the address of the second pointer.
	//This concept is sometimes termed as Double Pointers.
	var v int = 892

	// taking a pointer
	// of integer type
	var pt1 *int = &v

	// taking pointer to
	// pointer to pt1
	// storing the address
	// of pt1 into pt2
	var pt2 **int = &pt1

	fmt.Println("The Value of Variable v is = ", v)

	// changing the value of v by assigning
	// the new value to the pointer pt1
	*pt1 = 400

	fmt.Println("Value stored in v after changing pt1 = ", v)

	// changing the value of v by assigning
	// the new value to the pointer pt2
	**pt2 = 150

	fmt.Println("Value stored in v after changing pt2 = ", v)
}
