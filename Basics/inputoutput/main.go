package main

import (
	"fmt"
)

// global variable
var url = "https://frontendmaster.com"

func init() {
	fmt.Println("B")
}

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func main() {
	stateTax, cityTax := calculateTax(100)
	fmt.Println("state Tax:", stateTax, "city Tax:", cityTax)

	// Fuction scoped variables
	// const pi float32 = 3.14
	PrintData()
	// fmt.Println(data.Countries)
}
