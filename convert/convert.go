package convert

import (
	"bufio"
	"fmt"
	"main/convert/langs"
	"main/util"
	"os"
	"strings"
)

// Define a map of functions used for the parsing
var convertFunctions = map[string]func(){
	"go":     langs.CreateGolang,
	"csharp": langs.CreateCSharp,
}

var reader = bufio.NewReader(os.Stdin)

// Execute the parsing of the json using the specified language
func execute(lang string, tname string, strjson string) {
	if util.IsValidJson(strjson) {
		langs.Str = strjson
		langs.ObjectName = strings.Trim(tname, " ")
		convertFunctions[lang]()
	} else {
		invalidInput(strjson, tname, HandlePath)
	}
}

// handle lang and call respective function
func HandleJson(lang string, tname string) {
	strjson, err := util.GetInput("\npaste the json here: ")

	if err != nil {
		HandleJson(lang, tname)
	}

	execute(lang, tname, strjson)
}

// Try to open the file and execute the parsing JSON to struct
func HandlePath(lang string, tname string) {
	// file, _ := os.ReadFile("test.txt")
	// strjson := string(file)
	// execute(lang, tname, strjson)

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
