package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
)

// for get and post we have shortcut methods
// for rest we dont have shortcut methods

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

// func makeGetRequest() {
//     fmt.Println("GET Request")

//     res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
//     HandleError(err)
//     defer res.Body.Close()

//     var todo ToDo
//     err = json.NewDecoder(res.Body).Decode(&todo)
//     HandleError(err)

//     fmt.Println("Todo: ", todo)
// }

// func makePostRequest() {
//     fmt.Println("POST Request")

//     todo := ToDo{
//         UserId:    1,
//         Title:     "Dollar Sign One Time",
//         Completed: false,
//     }

//     jsonData, err := json.Marshal(todo)
//     HandleError(err)

//     jsonString := string(jsonData)
//     jsonReader := strings.NewReader(jsonString)

//     myURL := "https://jsonplaceholder.typicode.com/todos"

//     res, err := http.Post(myURL, "application/json", jsonReader)
//     HandleError(err)
//     defer res.Body.Close()

//     fmt.Println("Status Code: ", res.StatusCode)
// }


// we have to create a get reqest it doesent qutomatically exist
func makeGetRequest() {
    fmt.Println("GET Request")

    myURL := "https://jsonplaceholder.typicode.com/todos/1"
    req, err := http.NewRequest(http.MethodGet, myURL, nil)
    HandleError(err)

    client := http.Client{}
    res, err := client.Do(req)
    HandleError(err)
    defer res.Body.Close()

    var todo ToDo
    err = json.NewDecoder(res.Body).Decode(&todo)
    HandleError(err)

    fmt.Println("Todo: ", todo)
}

// we have to create a post reqest it doesent qutomatically exist
func makePostRequest() {
    fmt.Println("POST Request")

    todo := ToDo{
        UserId:    1,
        Title:     "Dollar Sign One Time",
        Completed: false,
    }

    jsonData, err := json.Marshal(todo)
    HandleError(err)
    jsonReader := strings.NewReader(string(jsonData))

    myURL := "https://jsonplaceholder.typicode.com/todos"
    req, err := http.NewRequest(http.MethodPost, myURL, jsonReader)
    HandleError(err)
    req.Header.Set("Content-Type", "application/json")

    client := http.Client{}
    res, err := client.Do(req)
    HandleError(err)
    defer res.Body.Close()

    fmt.Println("Status Code: ", res.StatusCode)
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

}

// Function to make a PATCH request it doesent qutomatically exist
func makePatchRequest() {
    updatedData := map[string]interface{}{
        "title": "Todo One Updated",
    }

    jsonData, err := json.Marshal(updatedData)
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

func HandleError(err error) {
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
}

// ### Concepts Explained

// #### 1. **Goroutines**
// - **Definition**: Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions.
// - **Example**: Imagine you’re cooking dinner while also watching a movie. You’re able to do both tasks simultaneously. In Go, each task (like cooking or watching) can run in its own goroutine.

// #### 2. **The `sync.WaitGroup`**
// - **Definition**: A `WaitGroup` is used to wait for a collection of goroutines to finish executing. You can think of it as a way to synchronize the completion of multiple tasks.
// - **Example**: If you have several people in a relay race, you want to wait until all runners have completed their laps before declaring the race finished. The `WaitGroup` keeps track of how many runners (goroutines) are left.

// ### Files Breakdown

// #### First File
// - **sayHello and sayHi functions**: These functions are meant to run concurrently. When `main()` starts them with `go`, they execute at the same time.
// - **Time Sleep**: The program uses `time.Sleep` to give the main function a little pause to let the goroutines complete their tasks before exiting. Without this, the program might terminate before `sayHello` or `sayHi` finish executing.

// #### Second File
// - **worker function**: Each worker represents a task that runs in its own goroutine. The `WaitGroup` is used to signal when each worker is done.
// - **Loop in main()**: The loop starts multiple workers and tracks them with the `WaitGroup`, ensuring the main function waits for all workers to complete before printing the final message.

// ### What Are Channels in GoLang?

// - **Definition**: Channels are a way for goroutines to communicate with each other and synchronize their execution. They act like pipelines where you can send and receive values between goroutines.
// - **Example**: Think of a walkie-talkie system where two people can communicate. One person speaks (sends data) while the other listens (receives data). In Go, when one goroutine sends a message through a channel, another goroutine can receive it, allowing for efficient communication.

// ### Real-World Example of Channels

// Imagine a restaurant where chefs and waitstaff work together:
// - **Channel**: The order pad (like a channel) where waitstaff write down orders for the chefs.
// - **Sending an Order**: When the waitstaff takes an order, they write it on the pad and hand it to the chef (sending data through the channel).
// - **Chef Receives Order**: The chef looks at the pad, sees the order, and starts cooking (receiving data).

// This way, the waitstaff can take multiple orders while the chef prepares them, leading to efficient operations in the restaurant!

// ### Summary

// - **Goroutines** allow concurrent execution.
// - **WaitGroups** synchronize the completion of multiple tasks.
// - **Channels** facilitate communication between goroutines, enabling efficient data sharing and synchronization.