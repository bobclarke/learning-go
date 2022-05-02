package main

import "fmt"

func main() {
	var myString string
	myString = "Green"
	fmt.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	fmt.Println("myString is set to", myString)
}

func changeUsingPointer(s *string) {
	fmt.Println("Address of s is", s)
	fmt.Println("Value of s is", *s)
	*s = "Red"
}
