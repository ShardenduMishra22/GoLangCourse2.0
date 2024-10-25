package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("test.txt") // Renamed variable from error to err

	// if err != nil {
	// 	fmt.Println("There is an error:", err)
	// 	return // Add return to avoid further execution if file creation fails
	// }
	// defer file.Close()

	// fmt.Println("File Created Successfully")

	// StringContent := "Hi I am Shardendu Mishar and I am learning GoLang"
	// _, err = io.WriteString(file, StringContent) // Use err instead of undefined error variable
	// if err != nil {
	// 	fmt.Println("There is an error:", err)
	// 	return
	// }

	// Method - 1 => Won't Fail
	buffer := make([]byte, 1024)
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("There is an error:", err)
		return
	}
	defer file.Close()
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("There is an error:", err)
			return
		}
		fmt.Print(string(buffer[:n]))
	}

	// Method - 2 => Might Fail
	// content, err := os.ReadFile("test.txt") // Renamed variable from error to err
	// if err != nil {
	// 	fmt.Println("There is an error:", err)
	// 	return
	// }
	// fmt.Println("Content of the file is:", string(content))
}
