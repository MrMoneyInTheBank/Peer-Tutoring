package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz!")

	fmt.Printf("Enter your name: ") // using Printf so the next line begins on same console (can append \n on desire)
	var name string
	fmt.Scan(&name)

	fmt.Printf("Welcome to the quiz %v", name)

	fmt.Printf("\nEnter your age: ")
	var age uint
	fmt.Scan(&age)

	// comparing the age to 10 which will return a boolean
	// just checking if the user is old enough to play the game

	if age >= 10 {
		fmt.Println("You can play the game!")
	} else {
		fmt.Println("You can not play the game :(")
		return // exit the function if the user is not old enough to play
	}

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer1 string
	var answer2 string
	fmt.Scan(&answer1, &answer2)

	var answer string = answer1 + " " + answer2

	if answer == "RTX 3090" {
		fmt.Println("You are correct!")
	} else {
		fmt.Println("Sorry you got it wrong")
		fmt.Println("Hope you had fun!")
		return
	}

	fmt.Println("Hope you had fun!")
	return
}
