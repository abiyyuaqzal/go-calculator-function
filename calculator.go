package gocalculatorfunction

import "errors"

func Calcualate(a, b float64, operation string) (float64, error) {
	var result float64
	var errorResult error

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		errorResult = errors.New("invalid operation")
	}
	return result, errorResult
}
