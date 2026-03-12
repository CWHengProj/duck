package main

import (
	"fmt"
	"math/rand"
)

func getRandomResponse() string {
	var array = []string{
		"string1",
		"string2",
		"string3",
		"string4",
	}
	seed := rand.Perm(len(array))[0]
	return array[seed]
}

func clearScreen() {
	logger.Println("Screen has been cleared..")
	fmt.Print("\033[H\033[2J") //TODO: understand ansi escape sequences (this is basically ctrl + L)
	fmt.Println("Conversation cleared. Let's quack things up!")
}

func printANSI() {
	logger.Println("Exiting Application..")
	//TODO: retrieve this from external file so that it is easier for user customization
	fmt.Println("QUACKKKKKKK printing ANSI ART")
}
