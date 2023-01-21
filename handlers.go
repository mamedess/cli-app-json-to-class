package main

import (
	"bufio"
	"fmt"
	"strings"
)

var wInputMsg string = "wrong input pal, try again!"

// handle golang choice by validating the json and calling the specified converter function if valid
func handleInput(r *bufio.Reader, f func(string, *bufio.Reader)) {
	input, _ := getInput("\nalright, hit me with it, paste your json in here: ", r)

	isValidJson(input, r)

	fmt.Println("alright, the json is valid, we are working on it!")
	f(input, r)
}

// called when user input is unknown
func wrongInput() {
	fmt.Println()
	fmt.Println(wInputMsg)
}

// get user prompt and returns it as a string, and error if any
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
