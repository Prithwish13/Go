package basics

import (
	"fmt"
	"slices"
)

func slicesF() {

	// var numbers = []int{}

	// numbers1 := []int{1, 2, 3, 4, 5}
	// numbers2 := []int{6, 7, 8, 9, 10}

	// slice := make([]int,  5) // Create a slice with a capacity of 10 but no initial elements.
	a := [5]int{1, 2, 3, 4, 5} // Create an array with 5 elements.

	slice1 := a[1: 4]
	fmt.Println("initial cap", cap(slice1)) // Print the capacity of the slice
	fmt.Println("initial length", len(slice1)) // Print the length of the slice

	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)

	fmt.Println("after append cap", cap(slice1)) // Print the capacity after appending

	fmt.Println(slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println(sliceCopy)

	// mil slice
	// var nilSlice []int // Declare a nil slice

	for i, v := range slice1 {
		fmt.Println(i, v) // This will not print anything since nilSlice is empty
	}

	// Slice Equal methods
	// Check if two slices are equal

	sliceCopy[0] = 111

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}

	// multidimensional slices
	multiSlice := make([][]int, 3)

	for i := 0; i < 3; i++{
		innerLength := i + 1
		multiSlice[i] =  make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			multiSlice[i][j] = i + j
		}
	}

	fmt.Println(multiSlice)

	fmt.Println("The capacity of the slice is:", cap(slice1))
	fmt.Println("The length of the slice is:", len(slice1))
	
}