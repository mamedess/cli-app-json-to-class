package main

import (
	"bufio"
	"main/convert"
	"os"
)

var validLangs = [2]string{"go", "csharp"}
var reader = bufio.NewReader(os.Stdin)

func main() {
	start()
}

func start() {
	convert.HandlePath("go", "class test")

	// tname, _ := util.GetInput("\nclass name:")
	// lang, _ := util.GetInput("\nlanguage:")

	// if !isValidLang(lang, validLangs) {
	// 	fmt.Println("\ninvalid input, press 'enter' key to try again:")
	// 	reader.ReadLine()
	// 	start()
	// } else {
	// 	config, _ := util.GetInput("\n(p/path, j/json, u/url) ")

	// 	switch config {
	// 	case "p":
	// 		convert.HandlePath(lang, tname)
	// 	case "j":
	// 		convert.HandleJson(lang, tname)
	// 	case "u":
	// 		convert.HandleUrl(lang, tname)
	// 	default:
	// 		panic("wrong input")
	// 	}
	// }
}
