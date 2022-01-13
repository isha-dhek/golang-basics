// code for pointer to struct
package main

import "fmt"

// first create a structure
type Employee struct {

	// take variables with their type
	name  string
	empid int
}

func main() {

	emp := Employee{"Go", 493}

	// We take a address pointer to the struct
	pts := &emp

	fmt.Println(pts)
	pts.name = "Lang"

	fmt.Println(pts)

	//creating array size of a pointer
	arr := [6]int{100, 200, 300,
		400, 500, 600}

	var x int

	var p [4]*int

	// For loop to Assign the address
	for x = 0; x < len(p); x++ {

		p[x] = &arr[x]
	}

	for x = 0; x < len(p); x++ {

		fmt.Printf("Value of p[%d] = %d\n", x, *p[x])
	}

	//using len() function to find the length
	fmt.Println("Length of arr: ", len(arr))
	fmt.Println("Length of p: ", len(p))

}
