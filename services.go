package main

import (
	"fmt"
)

func getRandomResponse() string {
	//TODO: random string return
	var array = []string{
		"string1",
		"string2",
		"string3",
		"string4",
	}
	return array[0]
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") //TODO: understand ansi escape sequences (this is basically ctrl + L)
}

func printANSI() {
	//TODO: retrieve this from external file so that it is easier for user customization
	fmt.Println("QUACKKKKKKK printing ANSI ART")
}
