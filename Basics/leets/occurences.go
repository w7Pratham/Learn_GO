package main

import (
	"fmt"
)

func printMap(mapp map[string]int) {
	for key, value := range mapp {
		fmt.Println("key: ", key, ", value: ", value)
	}
}

func areOccurencesEqual(s string) bool {
	occurenceMap := make(map[string]int)

	for _, schar := range s {
		occurenceMap[string(schar)]++
	}

	if len(occurenceMap) == 0 {
		return true
	}

	firstCount := -1
	printMap(occurenceMap)

	for _, count := range occurenceMap {
		if firstCount == -1 {
			firstCount = count
		} else if firstCount != count {
			return false
		}
	}

	return true
}

func occurencesRunner() {
	fmt.Println(areOccurencesEqual("a"))
	fmt.Println(areOccurencesEqual("aabcc"))
}
