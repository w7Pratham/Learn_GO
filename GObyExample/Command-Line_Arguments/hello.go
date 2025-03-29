package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	if len(argsWithProg) < 3 {
		fmt.Println("There are not enough arguments.")
		fmt.Println("Atleast 3 arguments are required.")
		return
	}

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
