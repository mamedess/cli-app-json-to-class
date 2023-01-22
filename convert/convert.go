package convert

import (
	"bufio"
	"fmt"
	"main/convert/langs"
	"main/util"
	"os"
)

var convertFunctions = map[string]func(){
	"go":     langs.ConvertGolang,
	"csharp": langs.ConvertCSharp,
}

var reader = bufio.NewReader(os.Stdin)

func execute(lang string, strJson string) {
	if util.IsValidJson(strJson) {
		fmt.Println("alright, the json is valid, we are working on it!")
		langs.Str = strJson
		convertFunctions[lang]()
	} else {
		invalidInput(strJson, HandlePath)
	}
}

// handle lang and call respective function
func HandleJson(lang string) {
	strJson, err := util.GetInput("\nalright, hit me with it, paste the path here: ")

	if err != nil {
		HandleJson(lang)
	}

	execute(lang, strJson)
}

func HandlePath(lang string) {
	path, _ := util.GetInput("\nalright, hit me with it, paste the path here: ")
	file, err := os.ReadFile(path)

	if err != nil {
		panic("invalid file")
	}

	strJson := string(file)
	execute(lang, strJson)
}

func HandleUrl(lang string) {
	fmt.Printf("hey")
}

func invalidInput(lang string, f func(string)) {
	fmt.Println("pretty sure that's not valid json, press 'enter' key and try again!")
	reader.ReadLine()
	f(lang)
}
