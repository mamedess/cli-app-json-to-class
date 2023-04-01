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
		langs.Str = strjson
		langs.ObjectName = strings.Trim(tname, " ")
		convertFunctions[lang]()
	} else {
		invalidInput(strjson, tname, HandlePath)
	}
}

func HandlePath(lang string, tname string) {
	path, _ := util.GetInput("\npaste the path here: ")
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

func invalidInput(lang string, tname string, callback func(string, string)) {
	fmt.Println("invalid json, press 'enter' key and try again!")
	reader.ReadLine()
	callback(lang, tname)
}
