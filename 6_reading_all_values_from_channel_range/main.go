package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 2, 3}

	c := make(chan int)
	go sum(s, c)

	for sum := range c {
		fmt.Println(sum)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		time.Sleep(time.Second * 2)
		c <- sum
	}
	close(c)
}
