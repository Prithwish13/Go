package basics

import (
	"fmt"
	"math"
)

func basic() {
	var a, b int = 10, 20
	result := a + b

	fmt.Println("Addition:", result)

	const pi float64 = 22/ 7.0 //division between two integer will also be a integer, go will lean towards to floor value 
	fmt.Println("Value of PI:", pi)
	

	// Over flow example signed integer
	var maxInt int = 1<<63 - 1 // maximum value for int64
	fmt.Println("Max Int:", maxInt)
	maxInt += 1 // This will cause an overflow
	fmt.Println("After Overflow:", maxInt)

	// Over flow example unsigned integer
	var maxUint uint = 1<<64 - 1 // maximum value for uint64
	fmt.Println("Max Uint:", maxUint)
	maxUint += 1 // This will cause an overflow
	fmt.Println("After Overflow:", maxUint)

	// under flow example signed integer

	var smallFloat float64 = 1.0e-323
	fmt.Println("Small Float:", smallFloat)
	smallFloat  = smallFloat /math.MaxFloat64 // This will cause an underflow
	fmt.Println("After Underflow:", smallFloat)

}
