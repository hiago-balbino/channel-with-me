package main

import (
	"fmt"
	"time"
)

func main() {
	word := "channel"
	channelLength := 4
	letters := make(chan rune, channelLength)

	go readWordToLetters(word, letters)

	for i := 0; i < channelLength; i++ {
		time.Sleep(time.Second * 2)
		fmt.Printf("letter read -> %c\n", <-letters)
	}
}

func readWordToLetters(word string, letters chan rune) {
	for _, letter := range word {
		letters <- letter
		fmt.Printf("letter written -> %c\n", letter)
	}
	close(letters)
}
