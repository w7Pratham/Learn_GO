package main

import (
	"fm-basics/imports"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	newTicket := imports.Ticket{
		ID:    123,
		Event: "FM masters",
	}

	newTicket.PrintEvent()

	fmt.Println(newTicket)
}
