package basics

import (
	"errors"
	"fmt"
)

func return_value() {
	q, r := divider(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	result, err := compare(5, 5)

	fmt.Println("Comparison Result:", result)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison is successful")
	}
	
}

func divider(a, b int)(quotient int, reminder int) {
	if b == 0 {
		return 0, 0 
	}
	quotient = a / b
	reminder = a % b
	return
}

func compare(a, b int)(string, error){
	if a == b {
		return "", errors.New("Values are equal unable to compare")
	} else if a < b {
		return "Less than", nil
	} else {
		return "Greater than", nil
	}
}