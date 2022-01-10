// array as an argument in the function
package main

import "fmt"

func myfun(a [6]int, size int) int {
	var k, val, r int

	for k = 0; k < size; k++ {
		val += a[k]
	}

	r = val / size
	return r
}

func main() {

	var arr = [6]int{45, 63, 44, 35}
	var result int

	result = myfun(arr, 6)
	fmt.Printf("Final result is: %d ", result)
}
