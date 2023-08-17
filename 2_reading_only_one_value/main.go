package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3}
	result := make(chan int)

	go sum(numbers, result)

	fmt.Println(<-result)
}

func sum(numbers []int, result chan int) {
	sum := 0
	for _, number := range numbers {
		sum += number
		time.Sleep(time.Second * 2)
		result <- sum
	}
}
