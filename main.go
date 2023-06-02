package main

import (
	"fmt"
	"os"
)

func main() {

	var num1, num2, result int
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator: ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Print("This is not a valid operator: ", operator)
		os.Exit(1)
	}

	fmt.Print("Answer: ", result)
}
