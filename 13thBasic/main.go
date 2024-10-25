package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
	Married bool
}

func main() {
	fmt.Println("Hello, World!")	

	var p1 Person;
	p1.FirstName = "Shardendu"
	p1.LastName = "Mishra"
	p1.Age = 30
	p1.Married = false
	
	fmt.Println(p1)
}