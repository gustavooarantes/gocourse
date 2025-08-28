package main

import "fmt"

// Go variables are mutable
func main() {
	// Declaring a variable
	var age int

	// Declaring and initializing a variable
	var firstName string = "John"

	// Using the := operator (the compiler will automatically infer the type)
	// It can only be used within functions (not available for global scope declarations)
	count := 10
	lastName := "Doe"

	// Default value for numbers: 0
	// Default value for boolean: false
	// Default value for strings: ""
	// Default value for pointers, slices, maps, functions and structs: nil
}
