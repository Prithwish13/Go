package basics

import "fmt"


func basics() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}

	// iterate over a slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range numbers {
		if value%2 == 0 {
			println(value, "at index", index, "is even")
		} else {
			fmt.Printf("%d at index %d is odd\n", value, index)
		}
	}

	// break and continue example
	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		println("Skipping 5")
	// 		continue // skip the rest of the loop when i is 5
	// 	}
	// 	if i == 8 {
	// 		println("Breaking at 8")
	// 		break // exit the loop when i is 8
	// 	}
	// 	println("Current number:", i)
	// }

	// rows := 5
	// outer loop for rows
	// for i := 1; i <= rows; i++ {
	// 	// inner loop for spaces
	// 	for j := rows; j > i; j-- {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for stars
	// 	for k := 1; k <= (2*i - 1); k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // move to the next line after each row
	// }


	for i:= range 10 {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}