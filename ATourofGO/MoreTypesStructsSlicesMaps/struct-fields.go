package main

import (
	"fmt"
)

// Struct fileds are accessed using a dot.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4          // Updating the struct field.
	fmt.Println(v.X) // Printing the struct filed.
}
