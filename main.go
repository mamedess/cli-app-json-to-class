package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	start(reader)
}

func start(r *bufio.Reader) {
	input, _ := getInput("Which lang are we working with?:", r)

	// a switch might not be the best way
	// may change it later
	// the ideia is to find the easiest way to
	// predict typos or just diff way of calling the language
	// have no ideia how to implement today
	switch input {
	case "golang":
	case "go":
		JsonToGolang(input)
	case "c#":
		fmt.Println("work in progress")
	case "typescript":
	case "ts":
		fmt.Println("work in progress")
	default:
		fmt.Println("not implemented")
		getInput("", r)
	}
}

func JsonToGolang(prompt string) string {
	var gs string = ""
	//read prompt

	//idk

	return gs
}
