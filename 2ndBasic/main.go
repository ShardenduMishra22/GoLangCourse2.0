package main

import "fmt"

func DivByZeroError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by zero error")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Hello, World!")
	// Greet()
	// fmt.Println(Adder(1,2,3,4))

	// checking the error
	result1, err := (DivByZeroError(1, 0))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}

	// normal working of the function
	result2, err := DivByZeroError(42, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
}

// func Greet(){
// 	fmt.Println("Hi Brother, how are you?")
// }

// the three function below are same

// func Adder(a int,b int,c int,d int) int {
// 	return a+b+c+d
// }

// func Adder(a,b,c,d int) (result int) {
//	return a+b+c+d
// }

// func Adder(a,b,c,d int) (result int) {
// 	result = a+b+c+d
// 	return
// }
