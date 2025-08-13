package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	const COMPANY_NAME = "Tech Solutions Inc."

	fmt.Println("Company Name:", COMPANY_NAME)

	employee := Employee{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	fmt.Println("Employee Details:")
	fmt.Println("First Name:", employee.FirstName)
	fmt.Println("Last Name:", employee.LastName)
	fmt.Println("Age:", employee.Age)

	// Using a different naming convention
	var emp Employee
	emp.FirstName = "Jane"
	emp.LastName = "Smith"
	emp.Age = 25

	fmt.Println("\nAnother Employee Details:")
	fmt.Println("First Name:", emp.FirstName)
	fmt.Println("Last Name:", emp.LastName)
	fmt.Println("Age:", emp.Age)

}