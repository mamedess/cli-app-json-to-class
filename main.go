package main

import (
	"bufio"
	"fmt"
	"main/convert"
	"main/util"
	"os"
)

var validLangs = [2]string{"go", "csharp"}
var reader = bufio.NewReader(os.Stdin)

func main() {
	start()
}

func start() {
	tname, _ := util.GetInput("\nclass name:")
	lang, _ := util.GetInput("\nlanguage:")

	if !isValidLang(lang, validLangs) {
		fmt.Println("\ninvalid input, press 'enter' key to try again:")
		reader.ReadLine()
		start()
	}

	config, _ := util.GetInput("\nWhat do you prefer? Pointing to a .txt/.json file or pasting the json here? (f/file, p/pasting, u/url) ")

	switch config {
	case "p":
		convert.HandleJson(lang, tname)
	case "f":
		convert.HandlePath(lang, tname)
	case "u":
		convert.HandleUrl(lang, tname)
	default:
		panic("wrong input")
	}
}

func isValidLang(str string, arr [2]string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
	}

	return false
}
