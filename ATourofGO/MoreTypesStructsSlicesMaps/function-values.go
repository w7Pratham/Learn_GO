package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, i, j float64) float64 {
	return fn(i, j)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 121))

	fmt.Println(compute(hypot, 5, 121))
	fmt.Println(compute(math.Pow, 3, 4))
}
