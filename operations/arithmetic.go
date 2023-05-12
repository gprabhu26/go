package main

// fmt package provides the function to print anything
import "fmt"

func main() {
	// declaring the float variables
	var number1, number2, answer float32
	fmt.Println("Enter the numbers on which you want to perform arithmetic operators.")

	// scanning the input by the user
	fmt.Println("Enter the first number.")
	fmt.Scanln(&number1)
	fmt.Println("Enter the second number.")
	fmt.Scanln(&number2)

	// declaring the string variable using the var keyword
	var operator string
	fmt.Println("Enter the operator you want to perform on two numbers.")

	// scanning the input by the user
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		answer = number1 + number2
		fmt.Println("The addition of", number1, "and", number2, "is", answer)
	case "-":
		answer = number1 - number2
		fmt.Println("The subtraction of", number1, "and", number2, "is", answer)
	case "*":
		answer = number1 * number2
		fmt.Println("The multiplication of", number1, "and", number2, "is", answer)
	case "/":
		answer = number1 / number2
		fmt.Println("The division of", number1, "and", number2, "is", answer)
	}
}
