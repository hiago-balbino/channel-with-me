package main

import (
	"fmt"
	"time"
)

func main() {
	word := "channel"
	letters := make(chan rune)

	go readWordToLetters(word, letters)

	fmt.Println(<-letters)
}

func readWordToLetters(word string, letters chan rune) {
	for range word {
		time.Sleep(time.Second * 2)
	}
}
