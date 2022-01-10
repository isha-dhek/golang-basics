// the nested structure
package main

import "fmt"

type Author struct {
	name   string
	branch string
	year   int
}

type HR struct {
	details Author
}

func main() {

	result := HR{

		details: Author{"Isha", "CSE", 2020},
	}

	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}
