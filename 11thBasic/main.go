package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	num := []int {12,14,16,18,20}
	for idx,val := range(num){
		fmt.Println("Index:",idx," Value:",val)
	}
}
