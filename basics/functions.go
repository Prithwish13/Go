package basics

import "fmt"


func functions() {

	result := add(3, 5)
	fmt.Println("Adding 3 and 5 gives:", result)

	// Anonymous function
	greet := func(){
		fmt.Println("This is an anonymous function")	
	}

	greet()


	operation := add

	fmt.Println("Using a function variable to add 10 and 20 gives:", operation(10, 20))

	fmt.Println("Using applyOperation to add 15 and 25 gives:", applyOperation(15, 25, operation))

	multiplierFunc := createMultiplier(3)

	fmt.Println("Multiplying 5 by 3 gives:", multiplierFunc(5))
}

func add(a, b int) int {
		return a + b
}
// functions that receives a function as an argument
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// functions that returns a function

func createMultiplier(factor int) func(int) int{
	return func(x int) int {
		return x * factor
	}
}