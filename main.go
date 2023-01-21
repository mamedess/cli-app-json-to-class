package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var config string

func main() {
	reader := bufio.NewReader(os.Stdin)

	start(reader)
}

func start(r *bufio.Reader) {
	getConfig(r)
	getLang(r)
}

// get wheather it is prefered to point to a file or to paste the json
func getConfig(r *bufio.Reader) {
	input, _ := getInput("\nWhat do you prefer? Pointing to a .txt/.json file or pasting the json here? (f/file, p/pasting) ", r)

	switch input {
	case "f":
		fmt.Println("\nnot implemented yet...")
		start(r)
	case "p":
		config = "p"
	default:
		wrongInput()
		start(r)
	}
}

// store wich lang will be used
func getLang(r *bufio.Reader) {
	input, _ := getInput("\nWhich lang are we working with?:", r)

	switch input {
	case "golang":
	case "go":
		handleInput(r, JsonToGolang)
	case "c#":
		handleInput(r, JsonToCSharp)
	default:
		wrongInput()
		getLang(r)
	}

}

// validate the json
func isValidJson(jsonPrompt string, r *bufio.Reader) {
	if !json.Valid([]byte(jsonPrompt)) {
		fmt.Println("\nwait a second, this is not valid JSON, press 'enter' key to try again")
		r.ReadLine()
		start(r)
	}
}
