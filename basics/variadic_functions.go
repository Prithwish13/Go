package basics

import "fmt"

func variadic() {

	sequence, result := sum(1, 1, 2, 3, 4, 5)

	fmt.Println("sequence", sequence, "Sum:", result)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence, result = sum(2, numbers...)

	fmt.Println("sequence", sequence, "Sum:", result)
	
}


func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return sequence, total
}