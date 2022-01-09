//The function that called with the varying number of arguments is known as variadic function.
//Or in other words, a user is allowed to pass zero or more arguments in the variadic function
//The variadic functions are generally used for string formatting.
//You can also pass multiple slice in the variadic function.
//You can not use variadic parameter as a return value, but you can return it as a slice.
// Go program to illustrate the
// concept of variadic function
package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(element ...string) string {
	return strings.Join(element, ",")
}

func main() {

	// pass a slice in variadic function
	element := []string{"Go", "Python", "C++"}
	fmt.Println(joinstr(element...))
}

//when we use variadic function : Variadic function is used when you want to pass a slice in a function.
//Variadic function is used when we don’t know the quantity of parameters.
