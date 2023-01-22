package convert

import (
	"bufio"
	"fmt"
	"main/convert/langs"
	"main/util"
	"os"
	"strings"
)

var convertFunctions = map[string]func(){
	"go":     langs.CreateGolang,
	"csharp": langs.CreateCSharp,
}

var reader = bufio.NewReader(os.Stdin)

func execute(lang string, tname string, strjson string) {
	if util.IsValidJson(strjson) {
		fmt.Println("\nalright, the json is valid, we are working on it!")
		langs.Str = strjson
		langs.Name = strings.Trim(tname, " ")
		convertFunctions[lang]()
	} else {
		invalidInput(strjson, tname, HandlePath)
	}
}

// handle lang and call respective function
func HandleJson(lang string, tname string) {
	strjson, err := util.GetInput("\nalright, hit me with it, paste the json here: ")

	if err != nil {
		HandleJson(lang, tname)
	}

	execute(lang, tname, strjson)
}

func HandlePath(lang string, tname string) {
	path, _ := util.GetInput("\nalright, hit me with it, paste the path here: ")
	file, err := os.ReadFile(path)

	if err != nil {
		panic("invalid file")
	}

	strjson := string(file)
	execute(lang, tname, strjson)
}

func HandleUrl(lang string, tname string) {
	fmt.Printf("hey")
}

func invalidInput(lang string, tname string, f func(string, string)) {
	fmt.Println("pretty sure that's not valid json, press 'enter' key and try again!")
	reader.ReadLine()
	f(lang, tname)
}
