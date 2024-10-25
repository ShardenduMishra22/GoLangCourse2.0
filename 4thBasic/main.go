package main

import "fmt"

func main() {
    // ===========================
    // Array Initialization
    // ===========================

	//#------------------------------------------------------------------------------#//
    
	// ===========================
	// Slice Initialization
	// ===========================


	// 1. Declaring an array with explicit length and assigning values
    var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array arr1:", arr1) // Output: [1 2 3 4 5]
	
	// 1. Declaring and initializing a slice with values directly
	var slc1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice slc1:", slc1) // Output: [1 2 3 4 5]
	
	//#------------------------------------------------------------------------------#//
	
    // 2. Using type inference with an array literal
    arr2 := [5]string{"A", "B", "C", "D", "E"}
    fmt.Println("Array arr2:", arr2) // Output: [A B C D E]
	
    // 2. Using type inference with a slice literal
    slc2 := []string{"X", "Y", "Z"}
    fmt.Println("Slice slc2:", slc2) // Output: [X Y Z]

	//#------------------------------------------------------------------------------#//

    // 3. Declaring an array with partial initialization
    arr3 := [5]float64{1.1, 2.2}
    fmt.Println("Array arr3:", arr3) // Output: [1.1 2.2 0 0 0]
	
    // 3. Creating a slice using make with specific length
    slc3 := make([]float64, 5)
    fmt.Println("Slice slc3:", slc3) // Output: [0 0 0 0 0]

	//#------------------------------------------------------------------------------#//

    // 4. Declaring an array without initialization (defaults to zero values)
    var arr4 [5]int
    fmt.Println("Array arr4:", arr4) // Output: [0 0 0 0 0]

    // 4. Creating a slice with make, specifying length and capacity
    slc4 := make([]int, 3, 5)
    fmt.Println("Slice slc4:", slc4) // Output: [0 0 0]
    fmt.Printf("Length of slc4: %d, Capacity of slc4: %d\n", len(slc4), cap(slc4))

	//#------------------------------------------------------------------------------#//

    // 5. Array with ellipsis (...) to let compiler infer length
    arr5 := [...]bool{true, false, true}
    fmt.Println("Array arr5:", arr5) // Output: [true false true]

    // 5. Slicing an array to create a slice
    baseArr := [5]int{10, 20, 30, 40, 50}
    slc5 := baseArr[1:4] // Creates a slice from index 1 to 3 (excludes 4)
    fmt.Println("Slice slc5 from baseArr:", slc5) // Output: [20 30 40]

	//#------------------------------------------------------------------------------#//
}
