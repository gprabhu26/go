package main

// fmt package provides the function to print anything
import "fmt"

func inputfromusermain() {

	// declaring the variable using the var keyword
	var numberFromUser int
	fmt.Println("Please enter the number you want to check that it is divisible by 2 or not:")

	// scanning the input by the user
	fmt.Scanln(&numberFromUser)
	// logic to check that number is divisible by 2 or not
	if numberFromUser%2 == 0 {
		fmt.Println(numberFromUser, "is divisible by 2")
	} else {
		fmt.Println(numberFromUser, "is not divisible by 2")
	}
}
