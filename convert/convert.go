package convert

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"main/convert/langs"
	"main/util"
	"net/http"
	"os"
	"strings"
	"time"
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
	path, _ := util.GetInput("\npaste the path here: ")
	file, err := os.ReadFile(path)

	if err != nil {
		panic("invalid file")
	}

	strjson := string(file)
	execute(lang, tname, strjson)
}

func HandlePathDebug() {
	file, _ := os.ReadFile("test.txt")
	strjson := string(file)
	execute("go", "test.txt", strjson)
}

func HandleUrlDebug() {
	url := "https://jsonplaceholder.typicode.com/users"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, _ := spaceClient.Do(req)

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)

	strjson := string(body)
	execute("go", "teste", strjson)
}

func HandleUrl(lang string, tname string) {
	url, _ := util.GetInput("\npaste the url here: ")

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, _ := spaceClient.Do(req)

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)

	strjson := string(body)
	execute(lang, tname, strjson)
}

func invalidInput(lang string, tname string, callback func(string, string)) {
	fmt.Println("invalid json, press 'enter' key and try again!")
	reader.ReadLine()
	callback(lang, tname)
}
