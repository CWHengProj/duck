package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var i string
	fmt.Println("quack - enter response here")
	for true {
		//TODO: start cli application here
		fmt.Scan(&i)
		//TODO: add support to check for "exit"
		//TODO: shift + enter for new line
		if strings.Split(i, " ")[0] == "exit" {
			fmt.Print("quack - gracefully exiting..")
			os.Exit(0)
		}
		//TODO: add option to clear the chat
		//TODO: add option to retain the chat history in logs
		fmt.Println(getRandomResponse())
	}
}
func init() {
	//TODO: init the constants
	// apikeyInit
	/// ttlInit
	//TODO: error handling for broken init
	//TODO: init logging service
	fmt.Println("appi started succesfully")
	//TODO: add init ascii here
}

