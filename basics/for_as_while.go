package basics

import "fmt"

func whileLoop() {
	i:=1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Print(i, " is even, ")
		} else {
			fmt.Print(i, " is odd, ")
		}
		i++
	}

	sum := 0

	// Using for as a while loop with break condition

	for  {
		 sum += 10
		fmt.Println("Current sum:", sum)
		if sum > 100 {
			break
		}
	}
}