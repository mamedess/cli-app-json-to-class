package main

import (
	"bufio"
	"fmt"
	"main/convert"
	"main/util"
	"os"
)

// var 5validLangs = [2]string{"go", "csharp"}
var reader = bufio.NewReader(os.Stdin)

func main() {
	// convert.HandlePathDebug()

	start()
}

func start() {
	tname, _ := util.GetInput("\nclass name:")
	// lang, _ := util.GetInput("\nlanguage:")
	lang := "go"

	if lang != "go" {
		fmt.Println("\ninvalid input, press 'enter' key to try again:")
		reader.ReadLine()
		start()
	} else {
		config, _ := util.GetInput("\n(p/path, u/url) ")

		switch config {
		case "p":
			convert.HandlePath(lang, tname)
		case "u":
			convert.HandleUrl(lang, tname)
		default:
			panic("wrong input")
		}
	}
	start()
}
