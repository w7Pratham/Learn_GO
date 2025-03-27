package main

import (
	"fmt"
)

// need to spend more time on defer understand
// how it affects the function and return of function

func a() {
	// A deferred function's arguments are evaluated
	// when the defer statment is evaluated.
	i := 0
	defer fmt.Printf("Printing i before : %d\n", i)
	i++

	fmt.Printf("Printing i later: %d\n", i)
	return
}

func b() {
	// Deferred function calls are executed in
	// Last In First Out order after the surrounding
	// functions returns.
	for i := 0; i < 4; i++ {
		defer fmt.Printf("The i is : %d\n", i)
	}
}

func c() (i int) {
	// Deferred functions may read and assign to
	// the returning function's named return values.
	fmt.Printf("The before defer value of i is : %d\n", i)
	defer func() { i++ }()
	fmt.Printf("The after defer value of i is : %d\n", i)
	return 1
}

func myUnderstanding() (i int) {
	// For better understanding of whatever happening
	// in above function c
	fmt.Printf("--> Function myUnderstanding starts. Initial value of i: %d\n", i)

	defer func() {
		fmt.Printf("--> INSIDE DEFERED FUNC (before i++). Current value of i: %d\n", i)
		i++
		fmt.Printf("--> INSIDE DEFERED FUNC (after i++). New value of i: %d\n", i)
	}()

	fmt.Printf("--> Function myUnderstanding continues after defer statment. Value of i: %d\n", i)

	fmt.Println("--> Function myUnderstanding executing 'return 1'")
	return 1
}

func main() {
	fmt.Println("\nRunning a()")
	a()

	fmt.Println("\nRunning b()")
	b()

	fmt.Println("\nRunning c()")
	fmt.Println(c())

	fmt.Println("Running myUnderstanding()")
	fmt.Printf("Returning value from function myUnderstanding : %d", myUnderstanding())
}
