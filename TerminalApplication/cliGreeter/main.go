package main

import (
	"flag"
	"fmt"
)

type GreeterConfig struct {
	Greeting string
	Name     string
	Times    int
}

func ArgumentChecker(nonFlagArgs []string, timeValue int, greeting string) GreeterConfig {
	var name string
	if len(nonFlagArgs) < 1 {
		name = "Arther"
	} else {
		// setting the last argument as name
		name = nonFlagArgs[0]
	}

	// cheking on times if less than 1
	if timeValue < 1 {
		timeValue = 1
	}

	return GreeterConfig{
		Greeting: greeting,
		Times:    timeValue,
		Name:     name,
	}
}

func (gc *GreeterConfig) GreetingPrinter() {
	for range gc.Times {
		fmt.Println(gc.Greeting, gc.Name)
	}
}

func main() {
	// defining flags
	greeting := flag.String("greeting", "Wellcome Master :", "a simple greeting")

	times := flag.Int("times", 1, "number of times to repeat")

	// parsing flags
	flag.Parse()

	// getting struct with proper values
	GreetValues := ArgumentChecker(flag.Args(), *times, *greeting)

	// passing strcut to printer to print the output
	GreetValues.GreetingPrinter()
}
