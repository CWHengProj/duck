package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

//go:embed art/duckShaded.ans
var duckArt []byte

var helpDescription = "shows all the shortcuts"
var clearDescription = "clears the screen, start fresh"
var exitDescription = "exits the current session"
var webSearchDescription = "uses your default browser to search for your query e.g. /websearch are ducks waterproof?"

func getRandomResponse() string {
	var array = []string{
		"Interesting question, let me think about it.",
		"Can you explain your question in more detail?",
		"What is the simplest change you can make now?",
		"What is one thing you need to know but you don't?",
		"Let's break this down. Explain to me what you already understand.",
		"Let's Quack this case right open. Tell me more.",
		"It is fine to slow down. Let's list out the requirements for your task.",
	}
	seed := rand.Perm(len(array))[0]
	return array[seed]
}

func clearScreen() {
	logger.Println("Screen has been cleared..")
	fmt.Print("\033[H\033[2J")
	fmt.Println("Conversation cleared. Let's quack things up!")
}

func webSearch(input string) {
	logger.Println("Searching the web with query: ", input)
	queryString := retrieveSearchURL(input)
	exec.Command("open", queryString).Run()
	//TODO: add search operators - https://duckduckgo.com/duckduckgo-help-pages/results/syntax
	fmt.Printf("\rPerforming Web Search...")
	time.Sleep(3000 * time.Millisecond)
	fmt.Printf("\rHope the web search helped!\n")

}

func retrieveSearchURL(input string) string {
	baseURL := "https://duckduckgo.com/?q=" //TODO: make this configurable
	words := strings.Split(input, " ")
	words = words[1:]
	query := ""
	for i, word := range words {
		if i == len(words)-1 {
			query += word
		} else {
			query += word + "+"
		}
	}
	return baseURL + query
}

func exit() {
	logger.Println("Graceful shutdown initiated..")
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func printANSI() {
	logger.Println("Printing ANSI art..")
	fmt.Print(string(duckArt))
}

func spinner(waitTime int) {
	for start := time.Now(); time.Since(start) < time.Duration(waitTime)*time.Second; {
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
	fmt.Println("/websearch:", webSearchDescription)

}
