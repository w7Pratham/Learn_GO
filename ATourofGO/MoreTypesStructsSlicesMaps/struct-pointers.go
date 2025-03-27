package main

import (
	"fmt"
)

// Struct fields can be accessed through a struct pointer.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v        // point to v
	p.X = 1e9      // updating struct field
	fmt.Println(v) // Printing v value
}
