package main

import (
	"bufio"
	"fmt"
	"main/convert"
	"main/util"
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
	input, _ := util.GetInput("\nWhat do you prefer? Pointing to a .txt/.json file or pasting the json here? (f/file, p/pasting) ", r)

	switch input {
	case "f":
		fmt.Println("\nnot implemented yet...")
		start(r)
	case "p":
		config = "p"
	default:
		util.WrongInput()
		start(r)
	}
}

// store wich lang will be used
func getLang(r *bufio.Reader) {
	input, _ := util.GetInput("\nWhich lang are we working with?:", r)

	switch input {
	case "golang":
	case "go":
		util.HandleInput(r, convert.JsonToGolang)
	case "c#":
		util.HandleInput(r, convert.JsonToCSharp)
	default:
		util.WrongInput()
		getLang(r)
	}

}
