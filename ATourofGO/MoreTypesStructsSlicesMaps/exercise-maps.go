package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)

	for _, word := range strings.Fields(s) {
		countMap[word]++
	}
	return countMap
}

func main() {
	// WordCount("This is the string!")
	wc.Test(WordCount)

}
