package main

import "fmt"

func main() {
//  example of a valid input
	process(10)

	//  example of a in-valid input
	process(-10)

}

func process(input int) {

	defer fmt.Println("DEfer 1")
	defer fmt.Println("DEfer 2")
	if input < 0 {
		fmt.Println("before panic")
		panic("Negative input not allowed")
	}
	fmt.Println("Processing input:", input)
}