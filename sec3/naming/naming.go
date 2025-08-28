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
	
	// Idiomatic conventions:
	// CamelCase for exported functions and method names, camelCase for internal.
	// Acronyms should be all-caps (ParseHTTP, parseURL).
	// lowercase for local variables and package names, camelCase for exported.
	// CamelCase for constants and structs.
	
	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
