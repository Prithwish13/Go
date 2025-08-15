package basics

import (
	"fmt"
	"math/rand"
	"time"
)




func guessGame() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)


	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100. Try to guess it!")
	fmt.Println("Can you guess that number? to exit the game, type '0'.")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == 0 {
			fmt.Println("Exiting the game. Goodbye!")
			break
		}

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You've guessed the number:", target)
			break
		}
	}


}