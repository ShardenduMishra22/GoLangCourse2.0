package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "mango,apple,banana"
	fmt.Println(strings.Split(str1, ","))

	str2 := "        Malum Hai Naa         "
	str2 = strings.TrimSpace(str2)
	fmt.Println(str2)

	str3 := "Hello Hello Hello Hello Hello"
	fmt.Println(strings.Replace(str3, "Hello", "Hi", 3))

	str4 := "Hello Hello Hello Hello"
	fmt.Println(strings.Count(str4, "Hello"))

}
