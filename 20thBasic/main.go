package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Learning about JSON in GoLang")

	person := Person{
		Name:    "John Doe",
		Age:     25,
		IsAdult: true,
	}

	// Marshalling / Encoding => GoLang Object to JSON
	jsonData, _ := json.Marshal(person)
	fmt.Println("The JSON Data is : ", string(jsonData))
	// fmt.Println("The JSON Data is : ", jsonData) //The JSON Data is :  [123 34 110 97 109 101 34 58 34 74 111 104 110 32 68 111 101 34 44 34 97 103 101 34 58 50 53 44 34 105 115 95 97 100 117 108 116 34 58 116 114 117 101 125]

	// Unmarshalling / Decoding => JSON to GoLang Object
	var newPerson Person
	_ = json.Unmarshal(jsonData, &newPerson)
	fmt.Println("The Person is : ", newPerson)
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}
