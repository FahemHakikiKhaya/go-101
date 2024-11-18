package main

import (
	"errors"
	"fmt"
)

func main() {
	
	// ==============================
	// SECTION 2: Functions
	// ==============================

	// Function call
	printMe()

	// Using a function with multiple return values
	var result, remainder, err = intDivision(3, 2)

	// Handling errors and results
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	}

	// Using a switch statement with results
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}



// ==============================
// SECTION 3: Helper Functions
// ==============================

// printMe prints a message.
func printMe() {
	fmt.Println("Function")
}

// intDivision performs integer division and returns result, remainder, and an error if any.
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, nil
}