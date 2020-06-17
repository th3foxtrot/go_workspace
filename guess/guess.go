package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: Make this a game where either the computer guesses the number or you
// 1. Make the program print out how many tries it took.
// 2. See if you can tell if the user is lying?

// "input size" = 100, 100/2 guesses before it gets to the answer.
// n = 1 to 100 O(n)
// We'll use binary search instead.
// n -> binary search log(n)

func main() {
	// Recieve input from the command line
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan()

	for {
		// Binary Search
		guess := (low + high) / 2
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Invalid response, try again.")
		}
	}
}
