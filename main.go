package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var i string
	fmt.Println("quack - enter response here (TODO intro art here)")
	for true {
		fmt.Scan(&i)
		//TODO: shift + enter for new line
		switch firstWord := strings.Split(i, " ")[0]; firstWord {
		case "exit":
			fmt.Println("quack - gracefully exiting..")
			os.Exit(0)
		case "clear":
			clearScreen()
		default:

			// Faux loading
			// TODO: using ANSI art for polish phase
			for i := 0; i <= 100; i++ {
				fmt.Printf("\rLoading... %d%%", i)
				time.Sleep(10 * time.Millisecond)
			}

			//TODO: add option to retain the chat history in logs
			fmt.Println(getRandomResponse())
		}
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
