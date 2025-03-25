package main

import (
	"fmt"
	"strconv"
)

func itoaTest(s string) {
	var convertedString string

	for index, value := range s {
		fmt.Println(strconv.Atoi("1"))
		convertedString += strconv.Itoa(int(value) - 96)
		fmt.Println("int of value is :", int(value), "converted string after", index, "is :", convertedString)
	}
}
