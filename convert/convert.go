package convert

import (
	"bufio"
	"fmt"
	"main/convert/langs"
	"main/util"
	"os"
	"regexp"
)

// variable with all the properties of the json
var jsonProp = []string{}
var convertFunctions = map[string]func(){"go": langs.ConvertGolang, "csharp": langs.ConvertCSharp}

var reader = bufio.NewReader(os.Stdin)

// handle lang and call respective function
func HandleJson(lang string) {
	strJson, err := util.GetInput("\nalright, hit me with it, paste the path here: ")

	if err != nil {
		HandleJson(lang)
	}

	if util.IsValidJson(strJson) {
		fmt.Println("alright, the json is valid, we are working on it!")
		convertFunctions[lang]()
	} else {
		invalidJson(lang, HandleJson)
	}
}

func HandlePath(lang string) {
	path, _ := util.GetInput("\nalright, hit me with it, paste the path here: ")
	file, err := os.Open(path)
	strJson := ""
	if err != nil {
		HandlePath(lang)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strJson = scanner.Text()
	}

	if util.IsValidJson(strJson) {
		fmt.Println("alright, the json is valid, we are working on it!")
		JsonToGolang(path)
	} else {
		invalidJson(lang, HandlePath)
	}
}

func invalidJson(lang string, f func(string)) {
	fmt.Println("pretty sure that's not valid json, press 'enter' key and try again!")
	reader.ReadLine()
	f(lang)
}

func getPropNames(prompt string) {
	//get a slice with the whole property and it's value in the format of a string
	pattern := getPattern("ck$")
	m := pattern.Split(prompt, -1)

	//get a slice with the names of the properties
	pattern = getPattern("ck$")
	n := pattern.Split(prompt, -1)

	for i := 0; i < len(m); i++ {
		jsonProp = append(jsonProp, n[i])
	}
}

func JsonToGolang(prompt string) {
	getPropNames(prompt)
	// sb := CreateBaseJson()
	for i := 0; i < len(jsonProp); i++ {
		// appendTo(sb)
	}
}

func JsonToCSharp(prompt string) {
	getPropNames(prompt)
	// sb := CreateBaseJson()
	for i := 0; i < len(jsonProp); i++ {
		// appendTo(sb)
	}
}

func getPattern(regex string) *regexp.Regexp {
	return regexp.MustCompile(regex)
}
