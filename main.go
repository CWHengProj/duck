package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var logger *log.Logger

func main() {

	fmt.Println("quack - enter response here (TODO intro art here)")
	reader := bufio.NewReader(os.Stdin)
	for {
		var input string
		fmt.Println("Quack! Ask me anything")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) 
		//TODO: shift + enter for new line
		switch words := strings.Split(input, " "); words[0] {
		case "exit":
			fmt.Println("quack - gracefully exiting..")
			os.Exit(0)
		case "clear":
			clearScreen()
		default:
			logger.Println("Loading reply..")
			// TODO: using ANSI art for polish phase
			for i := 0; i <= 100; i++ {
				fmt.Printf("\rLoading... %d%%", i)
				time.Sleep(10 * time.Millisecond) //TODO: get the real value from config
			}
			//TODO: add option to retain the chat history in chatLogs
			fmt.Println("\n Quack:", getRandomResponse())
		}
	}
}
func init() {
	//TODO: init the constants
	// apikeyInit
	/// ttlInit
	//TODO: error handling for broken init

	//TODO: init logging service
	logPath := "./duck.log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //TODO: actually understand this
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Started new conversation with Duck")

	printANSI()

}

//TODO: sigterm on ctrl + c to trigger graceful shutdown
