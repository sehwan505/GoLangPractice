package main

import (
	"fmt"
	"math/rand"
)

func inputValue() (int, error) {
	var a int
	fmt.Printf("I've picked a number between 1 and 100. Can you guess it?\n")
	_, err := fmt.Scanln(&a)
	if err != nil {
		return 0, err
	}
	return a, nil
}

func main() {
	for {
		n := rand.Intn(100) + 1
		input, err := inputValue()
		if err != nil {
			fmt.Print("Invalid input. Please enter a number.\n")
			return
		}
		if input == n {
			fmt.Printf("You've guessed it right!\n")
		} else {
			fmt.Printf("You've guessed it wrong. The number was %d\n", n)
		}
	}
}
