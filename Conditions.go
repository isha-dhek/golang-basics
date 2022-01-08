// use of if..else..if ladder
package main

import "fmt"

func main() {

	// taking a local variable
	var v1 int = 50

	// checking the condition
	if v1 == 20 {

		// if condition is true then
		// display the following */
		fmt.Printf("Value of v1 is 20\n")

	} else if v1 == 70 {

		fmt.Printf("Value of a is 70\n")

	} else if v1 == 90 {

		fmt.Printf("Value of a is 90\n")

	} else {

		// if none of the conditions is true
		fmt.Printf("Value not matching\n")
	}
}
