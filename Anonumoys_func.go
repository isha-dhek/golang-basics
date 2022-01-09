//An anonymous function is a function which doesn’t contain any name.also known as function literal.
//syntax
//func(parameter_list)(return_type){
// code..

// Use return statement if return_type are given
// if return_type is not given, then do not
// use return statement
//	return
//	}()

// Go program to illustrate
// use of an anonymous function
package main

import "fmt"

func main() {

	// Assigning an anonymous
	// function to a variable
	value := func() {
		fmt.Println("Welcome to Earth!")
	}
	value()

}

//You can also return an anonymous function from another function.
//You can also pass an anonymous function as an argument into other function.
//You can also pass arguments in the anonymous function
