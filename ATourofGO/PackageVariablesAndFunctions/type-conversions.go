package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    fmt.Println(x, y, z)


    num := 42
    float_num := float64(num)
    uint_num := uint(float_num)
    fmt.Println(num, float_num, uint_num)
}
