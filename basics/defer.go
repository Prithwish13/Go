package basics

import "fmt"

func defer_func() {

	process(10)
	fmt.Println("main function")
	
}


func process(i int) {
	defer fmt.Println("deferred call in process with value:", i)
	defer fmt.Println("first deferred call in process")
	defer fmt.Println("second deferred call in process")
	defer fmt.Println("third deferred call in process")
	i++
	fmt.Println("processing")
	fmt.Println("process function with value:", i)
}