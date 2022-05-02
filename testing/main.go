package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := divide(20, 0)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("The result is", r)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by 0")
	}

	result = x / y
	return result, nil
}
