package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var z int
	fmt.Print("Enter a number (1-5): ")
	fmt.Scan(&z)

	switch {
	case z > 10:
		fmt.Println("You chose option 1: Performing addition.")
	case z < 10:
		fmt.Println("You chose option 2: Performing subtraction.")
	case z == 10:
		fmt.Println("You chose option 3: Performing multiplication.")
	default:
		fmt.Println("You chose option 5: Exiting the program.")
	}
}
