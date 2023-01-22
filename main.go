package main

import (
	"main/convert"
	"main/util"
)

func main() {
	start()
}

func start() {
	lang, _ := util.GetInput("\nWhich lang are we working with? (try your to be precise :D):")
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
