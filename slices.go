// Slice is a composite data type similar to an array which is used to hold the elements of the same data type.
// The main difference between array and slice is that slice can vary in size dynamically but not an array.
// A slice contains three components:

// Pointer: The pointer is used to points to the first element of the array that is accessible through the slice. Here, it is not necessary that the pointed element is the first element of the array.
// Length: The length is the total number of elements present in the array.
// Capacity: The capacity represents the maximum size upto which it can expand.

//using make() function ....  make() function is used to create an empty slice. Here, empty slices are those slices that contain an empty array reference.
// Go program to illustrate how to create slices
// Using make function
package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {

	// Creating an array of size 7
	// and slice this array till 4
	// and return the reference of the slice
	// Using make function
	var my_slice_1 = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_1, len(my_slice_1), cap(my_slice_1))

	// Creating another array of size 7
	// and return the reference of the slice
	// Using make function
	var my_slice = make([]int, 3)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice, len(my_slice), cap(my_slice))

	//using an array... synatx: array_name[low:high]

	arr := [4]string{"welcome", "to", "Earth!"}

	// Creating slices from the given array
	var myslc = arr[1:2]
	myslc1 := arr[0:]

	fmt.Println("My slice : ", myslc)
	fmt.Println("My slice : ", myslc1)

	//using slice literal
	var slice_1 = []string{"hello", "naruto"}
	fmt.Println((slice_1))

	//sorting of slice
	sort.Strings(slice_1)
	fmt.Println("slice 1 after sorting: ", slice_1)

	// Creating multi-dimensional slice
	s1 := [][]int{{12, 34},
		{56, 47},
		{29, 40},
		{46, 78}}

	fmt.Println("slice is: ", s1)

	// Iterate slice
	// using range in for loop
	// without index i.e Using a blank identifier in for loop:
	for _, ele := range slice_1 {
		fmt.Printf("Element = %s\n", ele)
	}

	//copying one slice into another ... syntax: func copy(dst, src []Type) int
	slcc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	//var slcc2 []int
	slcc3 := make([]int, 5)
	//slcc4 := []int{78, 50, 67, 77}

	copy_1 := copy(slcc3, slcc1)
	fmt.Println("\nSlice:", slcc3)
	fmt.Println("Total number of elements copied:", copy_1)

	//comparing two slices
	sli_1 := []byte{'I', 'S', 'H', 'A', 'D'}
	sli_2 := []byte{'I', 'S', 'H', 'A', 'A'}
	res := bytes.Compare(sli_1, sli_2)

	if res == 0 {
		fmt.Println("!..Slices are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!")
	}

}
