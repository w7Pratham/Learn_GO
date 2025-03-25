package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]bool

func init() {
	Countries[0] = "India"
	Countries[1] = "Russia"
	Countries[2] = "USA"

	qty := len(Countries)

	fmt.Println(qty)

	fmt.Println("Countries saved")
}
