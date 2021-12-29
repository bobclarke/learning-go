package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for x := 0; x < 3; x++ {
			fmt.Printf("Sending %d\n", x)
			ch <- x
			time.Sleep(time.Second)
		}
	}()

	for y := 0; y < 3; y++ {
		val := <-ch
		fmt.Printf("Got %d\n", val)

	}

}
