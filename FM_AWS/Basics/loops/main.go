package main

import (
    "fmt"
    "slices"
)

func main() {
    // animals := [2]string{}

    // animals[0] = "dog"
    // animals[1] = "cat"

    animals := []string{
        "dog",
        "cat",
    }
    
    fmt.Println("Before - ", animals)

    animals = append(animals, "moose")
    
    fmt.Println("After APPEND - ", animals)

    animals = slices.Delete(animals, 0, 1)
    
    fmt.Println("After DELETE - ", animals)

    for i := 0; i < len(animals); i++ {
        fmt.Println("This is my animal %s", animals[i])
    }

    for index, value := range animals {
        fmt.Println("This is my index %d and this is my animal %s\n", index, value)
    }

    for value := range 10 {
        fmt.Println(value)
    }

    // While loop is used with for only ;)
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
