package main

import "fmt"

type Person struct {
	FirstNamr string
	LastName string
	Age int
	Married bool
}

func main() {
	fmt.Println("Hello, World!")	

	var p1 Person;
	p1.FirstNamr = "Shardendu"
	p1.LastName = "Mishra"
	p1.Age = 30
	p1.Married = false
	
	fmt.Println(p1)
}