package main

import "fmt"

func main() {
	for x := 1; x <= 20; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Printf("x is %d - fizzbuzz\n", x)
		} else if x%5 == 0 {
			fmt.Printf("x is %d - buzz\n", x)
		} else if x%3 == 0 {
			fmt.Printf("x is %d - fizz\n", x)
		} else {
			fmt.Printf("x is %d\n", x)
		}

	}

}
