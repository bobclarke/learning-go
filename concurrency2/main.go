package main

import (
	"fmt"

	"github.com/bobclarke/learning02/helpers"
)

const numPool = 10

func calculateValues(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber

}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValues(intChan)

	num := <-intChan
	fmt.Println(num)
}
