package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Hello, World!")

	// // Example of converting from int to float64
	// var myInt int = 42
	// var myFloat64 float64 = float64(myInt) // Convert int to float64
	// fmt.Printf("Int: %d, Float64: %f\n", myInt, myFloat64)

	// // Example of converting from float64 to int
	// var myNewFloat64 float64 = 42.99
	// var myNewInt int = int(myNewFloat64) // Convert float64 to int
	// fmt.Printf("Float64: %f, Int: %d\n", myNewFloat64, myNewInt)

	// Example of converting from string to int
	var myString string = "12345"
	myInt, err := strconv.Atoi(myString) // Convert string to int
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("String: %s, Int: %d\n", myString, myInt)
	}

	// Example of converting from int to string
	var myNewInt int = 67890
	myNewString := strconv.Itoa(myNewInt) // Convert int to string
	fmt.Printf("Int: %d, String: %s\n", myNewInt, myNewString)
}
