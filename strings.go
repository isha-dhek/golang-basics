// using for range loop
package main

import (
	"fmt"
	"strings"
)

func main() {

	// String as a range in the for loop
	for index, s := range "Isha" {

		fmt.Printf("The index number of %c is %d\n", s, index)
	}
	str1 := "Go"
	str2 := "Programming"
	str3 := "Go"

	// Checking the string are equal
	// or not using == operator
	result1 := str1 == str2
	result2 := str2 == str3

	fmt.Println("\nResult 1: ", result1)
	fmt.Println("Result 2: ", result2)

	// Checking the string are not equal
	// using != operator
	result3 := str1 != str2
	result4 := str2 != str3

	fmt.Println("Result 3: ", result3)
	fmt.Println("Result 4: ", result4)

	// comparing strings
	resul1 := "Go" > "Programming"
	fmt.Println("\nResult 5: ", resul1)

	resul2 := "Go" < "Programming"
	fmt.Println("Result 6: ", resul2)

	resul3 := "Go" == "Go"
	fmt.Println("Result 7: ", resul3)

	fmt.Println(strings.Compare("Go", "Programming"))
	fmt.Println(strings.Compare("Go", "Go"))

	// concatenate strings
	fmt.Println("\nNew string 1: ", str1+str2)

	result := str1 + str2 + "Language"
	fmt.Println("Result 8: ", result)

	// trim string
	strn1 := ".....Welcome To Earth....."
	r1 := strings.Trim(strn1, ".....")

	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 9: ", r1)

	st := "Welcome-To-Earth"
	re1 := strings.Split(st, "-")

	fmt.Println("\nResult 10: ", re1)

}
