package console

import (
	"os"
	"fmt"
	"bufio"
)


func getCmd(scanner *bufio.Scanner) (string, error) {
	fmt.Print("> "); scanner.Scan()
  	cmd := scanner.Text()
  	return cmd, scanner.Err()
}

func printErr(err error) {
	if err != nil {
		fmt.Println("Error reading from input: ", err)
	}
}

func Start(quit chan bool) {

	scanner := bufio.NewScanner(os.Stdin) 

	for ;; {

  		cmd, err := getCmd(scanner); printErr(err)
		
		if cmd == "quit" {
			quit <- true
			break
		}
  	}
}

