package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int = &num

	fmt.Println("num: ", num)
	fmt.Println("ptr: ", ptr)
	fmt.Println("Value inside the ptr (same as num): ", *ptr)
}
