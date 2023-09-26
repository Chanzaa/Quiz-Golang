package main

import (
	"fmt"
	"strings"
)

func main() {
	// input text
	text := "selamat pagi"
	// // split text kepada words
	words := strings.Split(text, "")

	for _, word := range words {

		fmt.Println(word)
	}

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}
	fmt.Println(counts)
}
