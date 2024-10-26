package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET POST PUT PATCH DELETE in GoLang")

	// url = https://jsonplaceholder.typicode.com

	// GET :=
	// makeGetRequest()

	// POST :=
	// makePostRequest()

	// PUT :=
	// makePutRequest()
	
	// Patch :=
	// makePatchRequest()

	// Delte :=
	// makeDeleteRequest()
}

// Sample Response
// {
//     "userId": 1,
//     "id": 1,
//     "title": "delectus aut autem",
//     "completed": false
// }

type ToDo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Function to make a PATCH request
func makePatchRequest() {
	updatedData := map[string]interface{}{
		"title": "Todo One Updated",
	}

	jsonData, err := json.Marshal(updatedData)
	HandleError(err)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodPatch, myURL, strings.NewReader(string(jsonData)))
	HandleError(err)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := http.Client{}
	res, err := client.Do(req)
	HandleError(err)
	defer res.Body.Close()

	data, _ := io.ReadAll(io.Reader(res.Body))
	fmt.Println("Response from PATCH:", string(data))
}

// we have to create a delete reqest it doesent qutomatically exist
func makeDeleteRequest() {
	var myURL = "https://jsonplaceholder.typicode.com/todos/1" // Corrected URL

	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	HandleError(err)

	client := http.Client{}
	res, err := client.Do(req)
	HandleError(err)
	defer res.Body.Close()

	fmt.Println("Delete Request Status Code:", res.StatusCode)
}

// we have to create a put reqest it doesent qutomatically exist
func makePutRequest() {
	todo := ToDo{
		UserId:    1,
		Id:        1,
		Title:     "delectus aut autem",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	HandleError(err)

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	HandleError(err)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := http.Client{}
	res, err := client.Do(req)
	HandleError(err)
	defer res.Body.Close()

	data, _ := io.ReadAll(io.Reader(res.Body))
	fmt.Println(string(data))
}

func makePostRequest() {
	fmt.Println("POST Request")

	todo := ToDo{
		UserId:    1,
		Title:     "Dollar Sign One Time",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	HandleError(err)

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myURL, "application/json", jsonReader)
	HandleError(err)
	defer res.Body.Close()

	fmt.Println("Status Code: ", res.StatusCode)
}

func makeGetRequest() {
	fmt.Println("GET Request")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	HandleError(err)
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		fmt.Println("Error: ", res.StatusCode)
		return
	}

	var todo ToDo
	err = json.NewDecoder(res.Body).Decode(&todo)
	HandleError(err)

	fmt.Println("Todo: ", todo)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
