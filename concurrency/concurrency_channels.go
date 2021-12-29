package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for x := 0; x < 5; x++ {
			fmt.Printf("Sending %d\n", x)
			ch <- x
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		_ = i // Ignore i
		val := <-ch
		fmt.Printf("Received %d\n", val)
	}

}
