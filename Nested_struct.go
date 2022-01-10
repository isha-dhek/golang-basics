// Golang program to illustrate
// the nested structure
package main

import "fmt"

// Creating structure
type Author struct {
	name   string
	branch string
	year   int
}

// Creating nested structure
type HR struct {

	// structure as a field
	details Author
}

func main() {

	// Initializing the fields
	// of the structure
	result := HR{

		details: Author{"Isha", "CSE", 2020},
	}

	// Display the values
	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}
