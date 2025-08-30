package main

import "fmt"

func main() {

	// Going through a simple range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Iteration over collection
	numbers := [] int {1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		// Using Printf to print a formatted string (can include placeholders)
		fmt.Printf("Index: %d, has value %d\n", index, value)
	}

	// Break statement -> breaks out of the loop and moves on to the next statement
	// Continue statement -> breaks out of the CURRENT iteration of the loop and moves on to the next iteration of the same loop
	
	for i := 1; i <= 10; i++ {
		
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Odd number: ", i)

		if i == 5 {
			break
		}
	}

	rows := 5

	//Outer loop
	for i:= 1; i <= rows; i++ {
		// First inner loop
		for j := 1; j <= rows - i; j++ {
			fmt.Print(" ")
		}
		// Second inner loop
		for k := 1; k <= 2 * i - 1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("We have a lift off!")
}
