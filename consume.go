package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {

    req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/todos/2", nil)
    if err != nil {
        log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
    }


    req.Header.Add("Content-type", "application/json; charset=utf-8")

 
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Failed to send HTTP request: %v", err)
         panic("Something went wrong!")

    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
        panic("Something went wrong!")

    }

    defer resp.Body.Close()

    data := make(map[string]interface{})

    json.Unmarshal(body, &data)

    for key, value := range data {
        fmt.Printf("%s: %v\n", key, value)
    }
}