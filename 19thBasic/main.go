package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning Web Request in GoLang")
	BaseURL := "https://jsonplaceholder.typicode.com"

	parsedURL, _ := url.Parse(BaseURL)
	// fmt.Printf("The Scheme is: %T",parsedURL)

	fmt.Println("The Scheme is : ", parsedURL.Scheme)
	fmt.Println("The Host is : ", parsedURL.Host)
	fmt.Println("The Path is : ", parsedURL.Path)
	fmt.Println("The RawQuery is : ", parsedURL.RawQuery)

	parsedURL.Path = "/todos/1"
	parsedURL.RawQuery = "userId=1"
	parsedURL.Scheme = "http"
	parsedURL.Host = "jsonplaceholder.typicode.com"

	fmt.Println("The Scheme is : ", parsedURL.Scheme)
	fmt.Println("The Host is : ", parsedURL.Host)
	fmt.Println("The Path is : ", parsedURL.Path)
	fmt.Println("The RawQuery is : ", parsedURL.RawQuery)
}
