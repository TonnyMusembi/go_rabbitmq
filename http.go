package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type MyData struct {
    // Define your data structure based on the API response JSON
    Key1 string `json:"key1"`
    Key2 int    `json:"key2"`
}

func main() {
    // URL of the API
    apiUrl := "https://api.example.com/data"

    // Make an HTTP GET request to the API
    resp, err := http.Get(apiUrl)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Parse the JSON data
    var data MyData
    if err := json.Unmarshal(body, &data); err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the parsed data
    fmt.Println("Key1:", data.Key1)
    fmt.Println("Key2:", data.Key2)
}
