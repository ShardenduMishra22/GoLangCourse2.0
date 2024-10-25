package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice is :", slc)

	fmt.Printf("1.) The length of the slice is : %d\n", len(slc))
	fmt.Printf("2.) The capacity of the slice is : %d\n", cap(slc))

	slc = append(slc, 4, 5, 10)

	fmt.Printf("3.) The length of the slice is : %d\n", len(slc))
	fmt.Printf("4.) The capacity of the slice is : %d\n", cap(slc))
}