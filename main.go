package main

import (
	"calculator/functions"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: calc <number1> <operator> <number2>")
		return
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number: ", os.Args[1])
		debug.PrintStack() // Print the stack trace
		return
	}
	operator := os.Args[2]
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid number: ", os.Args[3])
		debug.PrintStack() // Print the stack trace
		return
	}

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, functions.Add(num1, num2))
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, functions.Subtract(num1, num2))
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, functions.Multiply(num1, num2))
	case "/":
		result, err := functions.Divide(num1, num2)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
	default:
		fmt.Println("Unknown operator: ", operator)
	}
}
