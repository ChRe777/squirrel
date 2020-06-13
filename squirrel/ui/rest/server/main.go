package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

import (
	"github.com/mysheep/squirrel/ui/rest/console"
)

// -------------------------------------------------------------------------------------------------

const (
	port = "8080"
)

var (
	quit = make(chan bool)
)

// -------------------------------------------------------------------------------------------------

type RequestMsg struct {
	Id int
  	Expression string
}

type ResponeMsg struct {
	Id int
  	Expression string
}

// -------------------------------------------------------------------------------------------------

func main() {
  
  	registerHandlers() 
  
  	go startServer()
  
  	printMsg()
  
  	go console.Start(quit)
  	
  	<-quit	// wait until quit signal send
}

// -------------------------------------------------------------------------------------------------

func printMsg() {
  	fmt.Println("Server started at port:", port)
  	fmt.Println("Quit server with CTRL+C or type 'quit' an hit ENTER")
}

func startServer() {
	http.ListenAndServe(":"+port, nil) 
}
 
func registerHandlers() {
 	http.HandleFunc("/repl", replHandler)
}

func writeRespone(w http.ResponseWriter, msg ResponeMsg) {
  	json, _ := json.Marshal(msg)
  	w.Write(json)
}
  	
func replHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
    case "GET":
        w.WriteHeader(http.StatusOK)        
        msg := ResponeMsg{Id: 1, Expression: "get was called"}
        writeRespone(w, msg)
    case "POST":
        w.WriteHeader(http.StatusCreated)
         msg := ResponeMsg{Id: 1, Expression: "post was called"}
        writeRespone(w, msg)
    case "PUT":
        w.WriteHeader(http.StatusAccepted)
        w.Write([]byte(`{"message": "put called"}`))
    case "DELETE":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "delete called"}`))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
    

}
