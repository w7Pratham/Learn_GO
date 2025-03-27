package main

import (
	"fmt"
)

// this is just a tinker code please don't mind it
func main() {
	count := 0
	for {
		fmt.Printf("Count is %d\n", count)
		if count > 10 {
			break
		} else if count%2 == 0 {
			count += 2
			continue
		}
		// Will never reach here
		count++
	}
}
