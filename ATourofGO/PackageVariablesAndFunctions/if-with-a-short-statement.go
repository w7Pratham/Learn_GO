package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) (float64, float64) {
	if v := math.Pow(x, n); v < lim {
		return v, 0
	}

	// Can not return v here as variables declared by the
	// statement are only in scope until the of the if.
	return math.Pow(x, n), lim
}

func powPrinter(x, n, lim float64) {
	power, limit := pow(x, n, lim)
	fmt.Printf("Anser is : %f and limit is %f\n", power, limit)
}

func main() {
	powPrinter(3, 2, 10)
	powPrinter(3, 3, 20)
}
