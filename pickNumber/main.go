package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		n := rand.Intn(100) + 1
		var a int
		fmt.Printf("I've picked a number between 1 and 100. Can you guess it?\n")
		fmt.Scanf("%d", &a)

		if a == n {
			fmt.Printf("Congratulations! You've guessed the number!\n")
		} else {
			fmt.Printf("Sorry, the number was %d\n", n)
		}
	}
}
