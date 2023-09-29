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

    // Now add all the required request headers, mostly APIs secured with access tokens
    // req.Header.Add("Authorization", "Bearer "+token)
    // And today's application development, most of wire format is JSON.
    // if the wire format we wanted to be different, we could request but server has to support it.
    req.Header.Add("Content-type", "application/json; charset=utf-8")

    // we can build our custom http clients but here we are using deafult http client
    // when we use custom http client, we can send context object, timeouts etc..
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Failed to send HTTP request: %v", err)
    }

    // here we are reading the response body from the server
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
    }

    // Always close the response body
    defer resp.Body.Close()

    data := make(map[string]interface{})

    // this step de-serializes JSON data to our native Golang data
    json.Unmarshal(body, &data)

    for key, value := range data {
        fmt.Printf("%s: %v\n", key, value)
    }
}