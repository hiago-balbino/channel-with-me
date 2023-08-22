package main

import (
	"fmt"
	"time"
)

func main() {
	word := "chan"
	letters := make(chan rune)

	go readWordToLetters(word, letters)

	fmt.Println(<-letters)
	fmt.Println(<-letters)
	fmt.Println(<-letters)
	fmt.Println(<-letters)
	fmt.Println(<-letters) // deadlock happens here
}

func readWordToLetters(word string, letters chan rune) {
	for _, letter := range word {
		time.Sleep(time.Second * 2)
		letters <- letter
	}
}
