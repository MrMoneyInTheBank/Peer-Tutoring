package main

import "fmt"

// allows to output content to the console and accept user input

// func main() {
// 	fmt.Println("Hello World!")
// }

func main() {
	fmt.Println("Welcome to the quiz!")

	// data types in go
	// string
	// int
	// uint
	// float (32 bit and 64 bit)
	// bool

	// var name string = "Ansh"
	// variables are dynamically typed but strongly and therefore type can't be changed

	// there is also a way to declare variables without explicity declaring the type
	// This is the recommended way to declare variables

	// name := "Ansh"
	// age := 20
	// fmt.Printf("Hello %v, are you %v?", name, age)

	// variables can be embedded within strings an outputted the console with the printf
	// function. It is similar to f-strings in python.
	// %v for variables, %d for decimals, $f for floats and others

	// accepting user input
	var name string // just declaring the variable but not initializing it
	// collecting user input
	fmt.Scan(&name) // storing user input in the name variable
	fmt.Printf("Hello %v, welcome to the quiz!", name)
}
