package basics

import "fmt"

func arrays() {
	// var months = [12] string {
	// 	"January", "February", "March", "April",
	// 	"May", "June", "July", "August",
	// 	"September", "October", "November",
	// }

	var months = [12]int{} // If we do not initialize the array, it will be filled with zero values, nad for the string it will be empty strings.

	for i := 0; i < len(months); i++ {
		println(months[i])
	}

	var originalArray = [5]int{1, 2, 3, 4, 5}
	var copiedArray = originalArray // This creates a copy of the original array that will be deep copy.

	copiedArray[0] = 10 // Modifying the copied array does not affect the original array.
	println("Original Array:", originalArray[0]) // Output: 1
	println("Copied Array:", copiedArray[0])     // Output: 10
	println("Copied Array:", originalArray[0])     // Output: 1


	for _, value := range originalArray {
		println("Value:", value)
	}

	a, _ := someFunction() // This is a multiple assignment in Go, where we can assign values returned by a function to multiple variables.
	println("a:", a, "we don't want to use b") 
	
	println(len(copiedArray))// Output: a: 1

	// comparing arrays
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}

	print("Are arr1 and arr2 equal? \n", arr1 == arr2) // Output: true

	// Multidimensional  array

	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	} 

	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			println("Matrix[", i, "][", j, "] =", matrix[i][j])
		}
	}

	// coping the address of the array by using the pointer
	var anotherArray *[5]int
	anotherArray = &originalArray
	anotherArray[0] = 100 // Modifying the array through the pointer will affect the original array.

	fmt.Println("value of originalArray:", originalArray)
	fmt.Println("value of anotherArray:", *anotherArray)

}

	func someFunction() (int, int) {
		return 1, 2
	}