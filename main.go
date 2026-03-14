package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var logger *log.Logger

func main() {

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdown
		fmt.Println()
		exit()
	}()

	printANSI()
	fmt.Println("What's on your mind?")
	//TODO: tips section to remind user of shortcuts?
	fmt.Println("(/help for tips)")
	reader := bufio.NewReader(os.Stdin)
	for {
		var input string
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		logger.Println("User input: ", input)
		//TODO: shift + enter for new line
		switch words := strings.Split(input, " "); words[0] {
		case "/exit":
			exit()
		case "/help":
			help()
		case "/clear":
			clearScreen()
		default:
			waitTime := 5 //TODO: get from config instead
			spinner(waitTime)
			fmt.Println("\nQuack:", getRandomResponse())
			
		}
	}
}

func init() {
	//TODO: phase 3 init the constants
	// apikeyInit
	/// ttlInit
	//TODO: error handling for broken init

	logPath := "./duck.log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //TODO: actually understand this
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Started new conversation with Duck")
}
