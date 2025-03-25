package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are in varibales/main.go")
	var myName string = "Prathamesh"
	myInt := 100
	myFloat := 10.0

	fmt.Printf("Hello my name is %s my int is %d my float is %f\n", myName, myInt, myFloat)
}
