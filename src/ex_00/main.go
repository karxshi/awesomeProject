package main

import (
	"errors"
	"fmt"
	"strconv"
)

const PLUS string = "+"
const MINUS string = "-"
const MULTIPLY string = "*"
const DIVIDER string = "/"

func main() {
	run()
}

func run() {
	println("Input left operand:")
	firstOperand := getInputString()
	println("Input operation:")
	operation := getInputString()
	println("Input right operand:")
	secondOperand := getInputString()
	calculate(firstOperand, operation, secondOperand)
}

func getInputString() string {
	var input string
	_, err := fmt.Scanf("%s", &input)
	errorHandler(err)
	return input
}

func calculate(firstOperand, operation, secondOperand string) {
	switch operation {
	case PLUS:
		printResult(operation, convertStringToFloat(firstOperand)+convertStringToFloat(secondOperand))
	case MINUS:
		printResult(operation, convertStringToFloat(firstOperand)-convertStringToFloat(secondOperand))
	case MULTIPLY:
		printResult(operation, convertStringToFloat(firstOperand)*convertStringToFloat(secondOperand))
	case DIVIDER:
		if convertStringToFloat(secondOperand) == 0 {
			err := errors.New("cannot divide by zero")
			errorHandler(err)
		} else {
			printResult(operation, convertStringToFloat(firstOperand)/convertStringToFloat(secondOperand))
		}
	default:
		err := errors.New("invalid operation")
		errorHandler(err)
	}
}

func printResult(operation string, result float64) {
	if operation == DIVIDER {
		fmt.Printf("%.3f", result)
	} else {
		fmt.Printf("%g", result)
	}
}

func convertStringToFloat(operand string) float64 {
	ret, err := strconv.ParseFloat(operand, 64)
	errorHandler(err)
	return ret
}

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("invalid input\ncause: %s\n", err.Error())
		run()
	}
}
