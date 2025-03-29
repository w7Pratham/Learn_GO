package main

import "fmt"

// An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
// flexible view into the elements of an array. In practice, slices are much more
// common than arrays.

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
