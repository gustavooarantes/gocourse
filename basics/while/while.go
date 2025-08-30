package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Using a for loop as a while would work (passing no condition yields an infinite loop, like while true does)
	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	// Other way to terminate the loop is simply adding a condition for it to break at a specific point
	sum := 0
	for {
		sum += 10
		fmt.Println("Sum:", sum)

		if sum >= 50 {
			break
		}
	}

	// Simple guessing game
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the guessing game!")
	fmt.Println("You should guess a number netween 1 and 100.")
	fmt.Println("Good luck!")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess > target {
			fmt.Println("The correct answer is lower than that! Try again.")
			continue
		} else if guess < target {
			fmt.Println("The correct answer is higher than that! Try again.")
			continue
		} else {
			fmt.Println("Congratulation! You guessed it!")
			break
		}
	}
}
