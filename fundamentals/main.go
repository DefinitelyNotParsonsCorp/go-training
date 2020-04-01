package main

import (
	"fmt"
)

func Factorial(x int) int {
	var result int
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return result
}

func NamedResultFactorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * NamedResultFactorial(x-1)
	}
	return
}

func Divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("Can't divide by zero")
	}
	return numerator / denominator, nil
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Factorial Result: %d\n", Factorial(10))
	fmt.Printf("NamedResultFactorial: %d\n", NamedResultFactorial(10))

	var x, y = 4, 2
	result, err := Divide(x, y)
	fmt.Printf("Divide Result: %d, %s\n", result, err)

	zero := 0
	result, err = Divide(4, zero)
	fmt.Printf("Divide Result: %d, %s\n", result, err)
}
