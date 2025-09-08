package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if we have enough arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator <operation> <num1> <num2>")
		fmt.Println("Operations: add, subtract, multiply, divide")
		os.Exit(1)
	}

	// Parse the operation and numbers
	operation := os.Args[1]
	num1, err1 := strconv.ParseFloat(os.Args[2], 64)
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	// Check for parsing errors
	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers")
		os.Exit(1)
	}

	// Perform the operation
	result := calculate(operation, num1, num2)

	// Print the result
	fmt.Println(result)
}

func calculate(operation string, num1, num2 float64) float64 {
	switch operation {
	case "add":
		return num1 + num2
	case "subtract":
		return num1 - num2
	case "multiply":
		return num1 * num2
	case "divide":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
		return num1 / num2
	default:
		fmt.Printf("Error: Unknown operation '%s'\n", operation)
		fmt.Println("Supported operations: add, subtract, multiply, divide")
		os.Exit(1)
		return 0
	}
}