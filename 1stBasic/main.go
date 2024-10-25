// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var Name string = "Shardendu Mishra"
	// fmt.Println(Name)

	// var Num1 int = 50
	// fmt.Println(Num1)

	// var Num2 float64 = 25.5
	// fmt.Println(Num2)

	// var OnOffButton bool = true
	// fmt.Println(OnOffButton)

	// learningpackagemanagement.Greet(Name)

	// var sentence string
	// fmt.Scan(&sentence)
	// fmt.Println(sentence)

	fmt.Printf("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	fmt.Println(sentence)
}