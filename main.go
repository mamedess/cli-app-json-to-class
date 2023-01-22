package main

import (
	"bufio"
	"fmt"
	"main/convert"
	"main/util"
	"os"
)

var validTypes = [4]string{"go", "golang", "c#", "csharp"}
var reader = bufio.NewReader(os.Stdin)

func main() {
	start()
}

func start() {
	lang, _ := util.GetInput("\nWhich lang are we working with? (try your to be precise :D):")

	if !isValidType(lang, validTypes) {
		fmt.Println("\nInvalid input, press 'enter' key to try again:")
		reader.ReadLine()
		start()
	}

	config, _ := util.GetInput("\nWhat do you prefer? Pointing to a .txt/.json file or pasting the json here? (f/file, p/pasting) ")

	switch config {
	case "p":
		convert.HandleJson(lang)
	case "f":
		convert.HandlePath(lang)
	default:
		util.WrongInput()
		start()
	}
}

func isValidType(str string, arr [4]string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
	}

	return false
}
