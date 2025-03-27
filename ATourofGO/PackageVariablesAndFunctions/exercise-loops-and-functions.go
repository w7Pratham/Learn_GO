package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 11; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Printf("our Sqrt(2) is: %g\n", Sqrt(2))
	fmt.Printf("math.Sqrt(2) is: %g", math.Sqrt(2))
}
