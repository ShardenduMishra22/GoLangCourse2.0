package main

import "fmt"

func main() {
	fmt.Println("Arary in GoLang")

	// var arr1 [5] int;
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 3

	// fmt.Println(arr1) // [1 2 3 0 0]

	// var arr2 [5] string
	// arr2[0] = "Hello"
	// arr2[1] = "World"
	// arr2[2] = "GoLang"

	// fmt.Printf("%q \n",arr2) // [Hello World GoLang "" ""]

	// array initialization
	var arr3 [5]int = [5] int {1,2,3,4,5}
	fmt.Println(arr3)

	// slice initialization
	var slc3 [5]int = [5] int {1,2,3,4,5}
	fmt.Println(slc3)
}