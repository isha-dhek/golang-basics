// concept of the promoted methods
package main

import "fmt"

// Structure
type details struct {

	// Fields of the
	// details structure
	name    string
	age     int
	gender  string
	psalary int
}

// Nested structure
type emp struct {
	post string
	eid  int
	details
}

// Method
func (d details) promotmethod(tsalary int) int {
	return d.psalary * tsalary
}

func main() {

	// Initializing the fields of
	// the employee structure
	values := emp{
		post: "HR",
		eid:  4567,
		details: details{

			name:    "Priya",
			age:     20,
			gender:  "Female",
			psalary: 890,
		},
	}

	// Promoted fields of the
	// employee structure
	fmt.Println("Name: ", values.name)
	fmt.Println("Age: ", values.age)
	fmt.Println("Gender: ", values.gender)
	fmt.Println("Per day salary: ", values.psalary)

	// Promoted method of the
	// employee structure
	fmt.Println("Total Salary: ", values.promotmethod(30))

	// Normal fields of
	// the employee structure
	fmt.Println("Post: ", values.post)
	fmt.Println("Employee id: ", values.eid)
}
