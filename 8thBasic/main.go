package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var num int
	fmt.Scan(&num)
	fmt.Println("You entered: ", num)

	if num > 18 {
		fmt.Println("You are an adult")
	} else if num == 18 {
		fmt.Println("Not Yet")
	} else {
		fmt.Println("You are a child")
	}
}
