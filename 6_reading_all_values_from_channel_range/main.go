package main

import (
	"fmt"
	"time"
)

func main() {
	word := "channel"
	letters := make(chan rune)

	go readWordToLetters(word, letters)

	for letter := range letters {
		fmt.Println(letter)
	}
}

func readWordToLetters(word string, letters chan rune) {
	for _, letter := range word {
		time.Sleep(time.Second * 2)
		letters <- letter
	}
	close(letters)
}
