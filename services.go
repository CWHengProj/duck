package main

import (
	"fmt"
	"math/rand"
	"os"
)

func getRandomResponse() string {
	var array = []string{
		"Interesting question, let me think about it.",
		"Can you explain your question in more detail?",
		"What is the simplest change you can make now?",
		"What is one thing you need to know but you don't?",
		"Let's break this down. Explain to me what you already understand.",
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
	logger.Println("Exiting Application..")
	//TODO: retrieve this from external file so that it is easier for user customization
	fmt.Println("QUACKKKKKKK printing ANSI ART")
}
