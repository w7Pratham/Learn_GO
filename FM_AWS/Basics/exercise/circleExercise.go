package main

import (
	"fmt"
)

const (
	PI = 3.14
)

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		Radius: radius,
	}
}

func (c Circle) Circumference() {
	circumference_of_circle := 2 * PI * c.Radius
	fmt.Println("My cicumference is ", circumference_of_circle)
	// return circumference_of_circle
}

func (c Circle) Area() {
	area_of_circle := PI * c.Radius * c.Radius
	fmt.Println("My area is ", area_of_circle)
}

func main() {
	// cuteCircle := Circle{
	// 	Radius: 1.5,
	// }

	cuteCircle := NewCircle(1.5)

	fmt.Println("I'm a circle my Radius is ", cuteCircle.Radius)
	cuteCircle.Circumference()
	cuteCircle.Area()
}
