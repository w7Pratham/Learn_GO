package main

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	var convertedString string
	var sum int

	for _, v := range s {
		convertedString += strconv.Itoa(int(v) - 96)
	}

	index := 0
	for index < k {
		for j := 0; j < len(convertedString); j++ {
			num, _ := strconv.Atoi(string(convertedString[j]))
			sum += num
		}
		convertedString = strconv.Itoa(sum)
		sum = 0
		index++
	}

	result, _ := strconv.Atoi(convertedString)
	return result
}

func LeetRunner() {
	// fmt.Println(strconv.Atoi("1"))
	// fmt.Println(strconv.Itoa(1))
	fmt.Println("getting Lucky?")
	fmt.Println(getLucky("iiii", 1))
	fmt.Println(strconv.Itoa(0))
	// fmt.Println(getLucky("leetcode", 2))
	// fmt.Println(getLucky("zbax", 2))
}
