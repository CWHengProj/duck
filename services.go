package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var helpDescription = "shows all the shortcuts"
var clearDescription = "clears the screen, start fresh"
var exitDescription = "exits the current session"

func getRandomResponse() string {
	var array = []string{
		"Interesting question, let me think about it.",
		"Can you explain your question in more detail?",
		"What is the simplest change you can make now?",
		"What is one thing you need to know but you don't?",
		"Let's break this down. Explain to me what you already understand.",
		"Let's Quack this case right open. Tell me more.",
	}
	seed := rand.Perm(len(array))[0]
	return array[seed]
}

func clearScreen() {
	logger.Println("Screen has been cleared..")
	fmt.Print("\033[H\033[2J") //TODO: understand ansi escape sequences (this is basically ctrl + L)
	fmt.Println("Conversation cleared. Let's quack things up!")
}

func exit() {
	logger.Println("Graceful shutdown initiated..")
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func printANSI() {
	logger.Println("Printing ANSI art..")
	data, err := os.ReadFile("art/duckShaded.ans")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}

func spinner(waitTime int) {
	for start := time.Now(); time.Since(start) < time.Duration(waitTime) * time.Second; { 
		fmt.Printf("\rThinking    \u2014")
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("\rThinking.   \\")
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("\rThinking..  |")
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("\rThinking... /")
		time.Sleep(250 * time.Millisecond)
	}
	//TODO: add option to retain the chat history in chatLogs
	fmt.Printf("\r             ")
}

func help() {
	//TODO: cool art or box here
	fmt.Println()
	fmt.Println("/help:", helpDescription)
	fmt.Println("/exit:", exitDescription)
	fmt.Println("/clear:", clearDescription)

}
