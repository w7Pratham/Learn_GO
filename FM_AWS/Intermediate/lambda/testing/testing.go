package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const hashedText = "$2a$10$5W22ztGESrxMK.MYnlwrJ.IRXfPmgx47VQzW/.kBBK4jfqGyU.aD."

const plainText = "Shivay"

func ValidatePassword(hashedPassword, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err == nil
}

func main() {
	if !ValidatePassword(hashedText, plainText) {
		fmt.Println("Password does not match")
	} else {
		fmt.Println("Password matches")
	}
}
