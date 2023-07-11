package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mymath "github.com/ivshil/calculator/src/helpers"
)

func main() {
	fmt.Println("Calculator \n input sum/substract/multiply/divide + 2 numbers divided by space")

	// Create a new scanner to read input from os.Stdin (console)
	scanner := bufio.NewScanner(os.Stdin)

	// Read the next line of input
	scanner.Scan()

	inputElements := strings.Split(scanner.Text(), " ")
	firstNumber, err := strconv.Atoi(inputElements[1])
	if err != nil {
		fmt.Printf("Error at firstNumber: %s", err)
	}
	secondNumber, err := strconv.Atoi(inputElements[2])
	if err != nil {
		fmt.Printf("Error at secondNumber: %s", err)
	}
	// strings.TrimRight(input, "\n")
	// fmt.Printf("%sTest", inputElements[2])

	// Get the input text from the scanner
	switch inputElements[0] {
	case "sum":
		fmt.Println(mymath.Sum(firstNumber, secondNumber))
	case "substract":
		fmt.Println(mymath.Substract(firstNumber, secondNumber))

	case "multiply":
		fmt.Println(mymath.Multiply(firstNumber, secondNumber))

	case "divide":
		fmt.Println(mymath.Divide(firstNumber, secondNumber))

	default:
		fmt.Println("ERROR, input is not  proper")
	}
}
