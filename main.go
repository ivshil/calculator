package main

import (
	"fmt"
	"src/mymath"
)

func main() {
	a := 12
	b := 27

	fmt.Println(mymath.Sum(a, b))
	fmt.Println(mymath.Substract(a, b))
	fmt.Println(mymath.multiply(a, b))
	fmt.Println(mymath.devide(a, b))
}

// fmt.Println("Hello, this is a basic console calculator, what do you want to do?")

// fmt.Println("To use the calculator,\n input a command (sum, substract, multiply or devide) with 2 numbers, separated by space:")

// Create a new scanner to read input from os.Stdin (console)
// scanner := bufio.NewScanner(os.Stdin)

// Read the next line of input
// scanner.Scan()

// inputElements := strings.Split(scanner.Text(), " ")

// firstNumber := inputElements[1]
// secondNumber := inputElements[2]

// Get the input text from the scanner
// switch scanner.Text() {
// case "sum":

// case "substract":

// case "multiply":

// case "devide":

// }

// fmt.Printf("Hello, %s!\n", input)
