package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
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
	fmt.Println("quack - enter response here (TODO intro art here)")
	reader := bufio.NewReader(os.Stdin)
	for {
		var input string
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		//TODO: shift + enter for new line
		switch words := strings.Split(input, " "); words[0] {
		case "exit":
			exit()
		case "clear":
			clearScreen()
		default:
			logger.Println("Loading reply..")
			// TODO: using ANSI art spinner for polish phase
			for i := 0; i <= 100; i++ {
				fmt.Printf("\rLoading... %d%%", i)
				time.Sleep(10 * time.Millisecond) //TODO: get the real value from config
			}
			//TODO: add option to retain the chat history in chatLogs
			fmt.Println("\nQuack:", getRandomResponse())
		}
	}
}
func init() {
	//TODO: init the constants
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
