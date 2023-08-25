package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	letters := make(chan string)
	numbers := make(chan int)

	go populateChannels(letters, numbers)

	// To avoid selecting the channel in the select-case and not print the default value,
	// you can set the channel to nil when !ok
	// `letters = nil and numbers = nil`
	for {
		select {
		case letter, ok := <-letters:
			if ok {
				fmt.Println(letter)
			}
		case number, ok := <-numbers:
			if ok {
				fmt.Println(number)
			}
		case <-ctx.Done():
			fmt.Println("timeout reached")
			return
		}
	}
}

func populateChannels(letters chan string, numbers chan int) {
	defer func() {
		close(letters)
		close(numbers)
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for _, letter := range []string{"A", "B", "C"} {
			time.Sleep(1 * time.Second)
			letters <- letter
		}
	}()

	go func() {
		defer wg.Done()

		for _, number := range []int{1, 2, 3, 4} {
			time.Sleep(1 * time.Second)
			numbers <- number
		}
	}()

	wg.Wait()
}
