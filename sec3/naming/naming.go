package main

import "fmt"

type Employee struct {
	firstName string
	lastName string
	age int
}

func main() {
	// Interface names: one-method interfaces are named by the method name
	// plus an -er suffix or similar modification to construct an agent noun.
	
	// The convention in Go is to use MixedCaps or mixedCaps rather than
	// underscores to write multiword names.
	
	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
