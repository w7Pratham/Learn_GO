package main

import "fmt"

// Slices are like references to arrays
// A slice does not store any data, it just describes a section of an underlying
// array.

// Changing the elements of a slice modifies the corresponding elements of its
// underlying array.

// Other slices that share the same underlying array will see those changes.

func main() {
	names := [4]string{
		"Zatch",
		"Zeno",
		"Gin",
		"Bankai",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	c := names[:]
	fmt.Println(c)
	c = append(c, "Kira")
	c[1] = "Zeno"
	fmt.Println(c)
	fmt.Println(names)
}
