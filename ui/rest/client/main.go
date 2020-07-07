package main

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	URL = "http://localhost:8080/repl"
)

// -------------------------------------------------------------------------------------------------

type RequestMsg struct {
	Id int
  	Expression string
}

// -------------------------------------------------------------------------------------------------

func createReq(jsonStr []byte) (*http.Request, error){

   	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    
    return req, err
}

func main() {
    
    reqMsg := RequestMsg{Id:1, Expression: "(cdr '(1 2 3))"}
  	jsonStr, _ := json.Marshal(reqMsg)
  	  
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