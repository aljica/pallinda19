package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	wordsInString := strings.Fields(s)
	for i := 0; i < len(wordsInString); i++ {
		var word = wordsInString[i]
		if val, ok := words[word]; ok {
			words[word] = val + 1
			continue
		}
		words[word] = 1
	}
	return words
}

func main() {
	wc.Test(WordCount)
}
