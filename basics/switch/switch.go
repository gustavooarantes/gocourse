package main

import "fmt"

func main() {

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
		case int:
			fmt.Println("It's an integer!")
		case float64:
			fmt.Println("It's a float!")
		case string:
			fmt.Println("It's a string!")
		default:
			fmt.Println("IDK")
	}
}
