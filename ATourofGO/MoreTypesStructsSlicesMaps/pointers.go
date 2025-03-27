package main

import (
	"fmt"
)

// The type *T is a pointer to a T value. Its zero value is nil.
// The & operator generated a pointer to its operand.
// The * operator dontes the pointer's underlying value.
// This is known is "dereferencing" or "indirecting".

func main() {
	i, j := 42, 2701

	p := &i                                                // point to i
	fmt.Println("Reading i thorugh the pointer *p : ", *p) // read i through the pointer
	*p = 21                                                // set i through the pointer
	fmt.Println("The new value of i is :", i)              // see the new value of i

	p = &j                                    // point to j
	*p = *p / 37                              // divide j through the pointer
	fmt.Println("The new value of j is: ", j) // see the new value of j
}
