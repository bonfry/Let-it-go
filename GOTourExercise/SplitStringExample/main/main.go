package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsMap := make(map[string]int)

	for _, word := range words {

		if _, ok := wordsMap[word]; ok == true {
			wordsMap[word] += 1
		} else {
			wordsMap[word] = 1
		}
	}

	return wordsMap
}

func main() {
	wc.Test(WordCount)
}
