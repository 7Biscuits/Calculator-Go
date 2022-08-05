package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Println("Enter the first no. : ")
	fmt.Scan(&num1)

	fmt.Println("Enter the second no. : ")
	fmt.Scan(&num2)

	fmt.Println("Enter the operator : ")
	fmt.Scan(&operator)

	fmt.Println(calculate(num1, num2, operator))
}

func calculate(num1, num2 float64, operator string) (float64) {
	var result float64
	switch operator {
		case "+":
			result = num1 + num2
			break
		case "-":
			result = num1 - num2
			break
		case "*":
			result = num1 * num2
			break
		case "/":
			result = num1 / num2
			break
	}
	return result
}