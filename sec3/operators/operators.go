package main

import "fmt"

func main() {

	// Arithmetic Operators
	num1 := 1
	num2 := 4
	fmt.Println("num1 + num2 = ", num1 + num2)
	fmt.Println("num2 - num1 = ", num2 - num1)
	fmt. Println("num1 * num2 = ", num1 * num2)
	fmt.Println("num2 / num1 = ", num2 / num1)
	fmt.Println("num2 % num1 = ", num2 % num1)

	// Logical Operators: AND = &&, OR = ||, and NOT = !
	// Bitwise Operators: Bitwise AND = &, Bitwise OR = |, Bitwise XOR = ^,
	// Bit clear (AND NOT) = &^, Left shift = <<, and Right shift = >>
	// obs.: &^ clears the bits in the first operand wherever the second
	// operand has 1s, and shift operators require unsigned or integer types,
	// and the shit count must be an integer.
}
