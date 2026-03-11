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

var clearScreen = func() {

	fmt.Print("\033[H\033[2J")
	//TODO: proper ansi escape sequence to clear screen

}
