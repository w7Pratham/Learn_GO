package main

import (
	"fmt"
	"time"
)

// Define the struct
type UserProfile struct {
	UserID    int
	Email     string
	IsActive  bool
	LastLogin time.Time
}

// Function that MODIFIES a UserProfile via a pointer.
// It accepts a pointer (*UserProfile) so it can change the original data.
func activateUser(userPtr *UserProfile) {
	if userPtr == nil {
		fmt.Println("Cannot activate a nil user profile.")
		return
	}
	fmt.Printf("--> Activating user ID %d (%s). Current active status: %t\n",
		userPtr.UserID, userPtr.Email, userPtr.IsActive)

	userPtr.IsActive = true        // Modify the field via the pointer
	userPtr.LastLogin = time.Now() // Modify another field

	fmt.Printf("--> User ID %d is now active. Last login: %s\n",
		userPtr.UserID, userPtr.LastLogin.Format(time.RFC1123))
}

func main() {
	fmt.Println("--- Creating and activating a user ---")

	// Requirement: We need a pointer (*UserProfile) to pass to activateUser.
	// Solution: Use '&' directly before the struct literal.
	// This creates the UserProfile in memory and immediately gives us its address (pointer).

	newUserPtr := &UserProfile{
		UserID:   101,
		Email:    "test@example.com",
		IsActive: false, // Initial state
		// LastLogin will be the zero value for time.Time initially
	}

	// newUserPtr is now of type *UserProfile

	fmt.Printf("Before activation: User ID %d, Active: %t, Email: %s\n",
		newUserPtr.UserID, newUserPtr.IsActive, newUserPtr.Email)
	// Note: Go automatically dereferences pointers for field access (newUserPtr.UserID works).

	// Pass the pointer to the function that modifies the struct.
	activateUser(newUserPtr)

	fmt.Printf("After activation: User ID %d, Active: %t, Last Login: %s\n",
		newUserPtr.UserID, newUserPtr.IsActive, newUserPtr.LastLogin.Format(time.RFC1123))

	// We can see the original struct data (pointed to by newUserPtr) has been modified.

	fmt.Println("\n--- Efficiency Aspect (Conceptual) ---")
	// Imagine UserProfile had many more fields (e.g., Bio string, Preferences map, etc.)
	// type LargeProfile struct { UserID int; Bio string /* 1MB */; /* many other fields... */ }
	// func processProfile(lp *LargeProfile) { /* ... */ }

	// processProfile(&LargeProfile{UserID: 202, Bio: "..."}) // Pass pointer (e.g., 8 bytes)

	// vs.
	// func processProfileCopy(lp LargeProfile) { /* ... */ }
	// processProfileCopy(LargeProfile{UserID: 202, Bio: "..."}) // Pass value (copies entire struct, potentially megabytes)

	fmt.Println("If UserProfile were very large, passing a pointer (&UserProfile{...}) avoids copying the entire struct's data.")
}
