package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"plugin"
	"io/ioutil"
)

import (
	"github.com/mysheep/squirrel/ui/rest/server/console"
	"github.com/mysheep/squirrel/interfaces"
	"github.com/mysheep/squirrel/evaluator"
	"github.com/mysheep/squirrel/types"
)

// -------------------------------------------------------------------------------------------------

const (
	port 		= "8080"
	pluginPath 	= "../../../bin/"
	MaxBytes 	= 1024 * 1024			// 1 MByte
)

var (
	quit = make(chan bool)
	
	parser 	interfaces.Parser
	printer interfaces.Printer	
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

	ui := "lisp"
	version := "1.0.0"
	
	fileParser := getFileName(ui, "parser", version)
	filePrinter := getFileName(ui, "printer", version)

 	parser  = loadParserPlugin(fileParser)
 	printer = loadPrinterPlugin(filePrinter)
  
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

func getResult(printer interfaces.Printer, e *types.Cell) string {
	return string(printer.Sprint(e))
}

func eval(parser interfaces.Parser, printer interfaces.Printer, expStr string) string {
	exp := parser.Parse([]byte(expStr)); res := evaluator.Eval(exp, nil)	
	resStr := getResult(printer, res)
	return resStr	
}
  	
func getRequestMsg(w http.ResponseWriter, r *http.Request)	(string, error) {
  	
  	var requestMsg RequestMsg
  	
  	r.Body = http.MaxBytesReader(w, r.Body, MaxBytes)	// max. 1MB    
    bodyBytes, err := ioutil.ReadAll(r.Body)    
    
    if err != nil {
    	return "", err
    }
    
    err = json.Unmarshal(bodyBytes, &requestMsg) 
    return requestMsg.Expression, err
}
        
func replHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	
    case "GET":
        w.WriteHeader(http.StatusOK)        
        msg := ResponeMsg{Id: 1, Expression: "get was called"}
        writeRespone(w, msg)
        
    case "POST":
        w.WriteHeader(http.StatusCreated)
         
        exp, err := getRequestMsg(w, r)
        
        if err != nil {
        	msg := ResponeMsg{Id: 1, Expression: err.Error()}
        	writeRespone(w, msg)
        } else {
  	      	responseExp := eval(parser, printer, exp)
  	      	msg := ResponeMsg{Id: 1, Expression: responseExp}
        	writeRespone(w, msg)    
        }
        
    
    case "PUT":
        w.WriteHeader(http.StatusAccepted)
        msg := ResponeMsg{Id: 1, Expression: "put was called"}
        writeRespone(w, msg)
        
    case "DELETE":
        w.WriteHeader(http.StatusOK)
        msg := ResponeMsg{Id: 1, Expression: "delete was called"}
        writeRespone(w, msg)
        
    default:
        w.WriteHeader(http.StatusNotFound)
        msg := ResponeMsg{Id: 1, Expression: "unknown method was called"}
        writeRespone(w, msg)
    }
    

}

// -------------------------------------------------------------------------------------------------

func getFileName(ui string, pluginName string, version string) string {
	file := pluginPath+pluginName+"_"+ui+"."+version+".so"
	return file
}

// loadParserPlugin loads the parser plugin
func loadParserPlugin(file string) interfaces.Parser {

	plugIn, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}

	parserSym, err := plugIn.Lookup("Parser")
	if err != nil {
		panic(err)
	}
	
	var parser interfaces.Parser
	parser, ok := parserSym.(interfaces.Parser)
	if !ok {
		fmt.Println("unexpected type from module symbol:" +file)
		os.Exit(1)
	}
	
	fmt.Printf("Plugin '%v' loaded. \n", file)
	
	return parser
}

// loadPrinterPlugin loads the printer plugin
func loadPrinterPlugin(file string) interfaces.Printer {

	plugIn, err := plugin.Open(file)	//(*Plugin, error)
	if err != nil {
		panic(err)
	}

	printerSym, err := plugIn.Lookup("Printer")	// func (p *Plugin) Lookup(symName string) (Symbol, error)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Plugin '%v' loaded. \n", file)
	
	return printerSym.(interfaces.Printer)
}
