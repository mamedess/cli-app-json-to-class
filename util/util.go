package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
)

var wInputMsg string = "wrong input pal, try again!"

// handle golang choice by validating the json and calling the specified converter function if valid
func HandleInput(r *bufio.Reader, f func(string, *bufio.Reader)) {
	input, _ := GetInput("\nalright, hit me with it, paste your json in here: ", r)

	if isValidJson(input, r) {
		fmt.Println("alright, the json is valid, we are working on it!")
		f(input, r)
	} else {
		r.ReadLine()
		// Start(r)
	}
}

// validate the json inside the start context
func isValidJson(jsonPrompt string, r *bufio.Reader) bool {
	isValid := !json.Valid([]byte(jsonPrompt))
	if isValid {
		fmt.Println("\nwait a second, this is not valid JSON, press 'enter' key to try again")
	}
	return isValid
}

// called when user input is unknown
func WrongInput() {
	fmt.Println()
	fmt.Println(wInputMsg)
}

// get user prompt and returns it as a string, and error if any
func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
