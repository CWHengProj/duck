package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var i string
	fmt.Println("quack - enter response here")
	for true {
		fmt.Scan(&i)
		//TODO: concurrent checks for multiple things? graceful exit?
		//TODO: shift + enter for new line
		if strings.Split(i, " ")[0] == "exit" {
			fmt.Println("quack - gracefully exiting..")
			os.Exit(0)
		}
		if strings.Split(i, " ")[0] == "clear" {
			fmt.Println("clearing chat")
			clearScreen
			fmt.Printf("\r"
			//TODO: add option to clear the chat
			os.Exit(0)
		}
		// Faux loading
		// TODO: using ANSI for polish phase
		for i := 0; i <= 100; i++ {
			fmt.Printf("\rLoading... %d%%", i)
			time.Sleep(50 * time.Millisecond)
		}

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

var clearScreen = func() {
    cmd := exec.Command(`printf '\33c\e[3J'`) // clears the scrollback buffer as well as the screen.
    cmd.Stdout = os.Stdout
    cmd.Run()
}
