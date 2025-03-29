package main

import (
	"fmt"
)

type Vehicle interface {
	Move(distance float64)
	ReportStatus() string
}

type Car struct {
	Make     string
	Model    string
	Fuel     float64
	Position float64
}

type Bicycle struct {
	Brand    string
	Position float64
}

func (c *Car) Move(distance float64) {
	fuelCost := distance * 0.1
	if fuelCost > c.Fuel {
		c.Fuel = 0
		c.Position = c.Fuel / 0.1
		fmt.Println("Insufficient fuel to complete the distance!")
		fmt.Println("Current Postion is :", c.Position)
		return
	}
	c.Fuel -= fuelCost
	c.Position += distance
	// fmt.Println("Total distnace travelled successfully: ", c.Position)
	return
}

func (c Car) ReportStatus() string {
	return fmt.Sprintf("Car: %s %s, Position: %.1f, Fuel: %.1f", c.Make, c.Model, c.Position, c.Fuel)
}

func (b *Bicycle) Move(distance float64) {
	if distance < 0 {
		fmt.Println("Distance can not be negative.")
		distance = 0
	}
	b.Position += distance
	// fmt.Println("Total distnace travelled successfully: ", b.Position)
	return
}

func (b Bicycle) ReportStatus() string {
	return fmt.Sprintf("Bicycle: %s, Position: %.1f", b.Brand, b.Position)
}

func operateVehicle(v Vehicle, distance float64) {
	fmt.Println(v.ReportStatus())
	v.Move(distance)
	fmt.Println(v.ReportStatus())
	fmt.Println("-----------------")
	return
}

func main() {
	Harrier := Car{
		Make:     "TATA",
		Model:    "Harrier",
		Fuel:     100,
		Position: 0.7,
	}

	Mtb := Bicycle{
		Brand:    "Redbull",
		Position: 1.9,
	}

	operateVehicle(&Harrier, 70.7)
	operateVehicle(&Mtb, 20.1)
}
