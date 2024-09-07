package main

import "errors"

func add(num1 float64, num2 float64) float64 {

	return num1 + num2
}

func subtract(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func divide(num1 float64, num2 float64) (float64, error) {

	if num2 == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return (num1/num2), nil
}