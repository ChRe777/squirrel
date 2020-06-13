package main

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
)

const (
	URL = "http://localhost:8080/repl"
)

func createReq(jsonStr []byte) (*http.Request, error){

   	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    
    return req, err
}

func main() {
    fmt.Println("URL:>", URL)

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
 	req, err := createReq(jsonStr)
 	if err != nil {
        panic(err)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}