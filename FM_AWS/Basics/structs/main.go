package main

import (
    "fmt"
)


type Person struct {
    Name    string
    Age     int
}

func NewPerson(name string, age int) Person {
    return Person{
        Name: name,
        Age:  age,
    }
}

func (p *Person) changeName( newName string ) {
    p.Name = newName
}


func main() {
    myPerson := Person{
        Name:   "shiv",
        Age:    26,
    }

    yourPerson := NewPerson("Pratham", 17)

    fmt.Printf("This is my person %+v", myPerson)
    fmt.Printf("\nThis is your person %+v", yourPerson)

    myPerson.Name = "Shivay"
    fmt.Printf("\nMy person shiv changed name to %+v", myPerson)

    yourPerson.changeName("Arther")

    fmt.Printf("\nYour Person Pratham changed name to %+v", yourPerson)
}
